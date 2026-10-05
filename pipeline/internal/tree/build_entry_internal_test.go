package tree

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func testDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(filepath.Dir(dir)); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(orig) })
	return filepath.Base(dir)
}

func TestBuildFileEntry_NewFile_CreatedAndModifiedAreNow(t *testing.T) {
	dir := testDir(t)
	path := filepath.Join(dir, "1_x.json")
	writeTestFile(t, path, `{}`)
	now := time.Now()

	parent, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}
	node := parent.Children[0]
	entry, err := buildFileEntry(node, nil, now)
	if err != nil {
		t.Fatal(err)
	}
	if !entry.Created.Equal(now) || !entry.Modified.Equal(now) {
		t.Errorf("expected created and modified to both be now, got %v / %v", entry.Created, entry.Modified)
	}
}

func TestBuildFileEntry_UnchangedContent_KeepsBothDates(t *testing.T) {
	dir := testDir(t)
	path := filepath.Join(dir, "1_x.json")
	writeTestFile(t, path, `{"same":true}`)

	parent, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}
	node := parent.Children[0]
	created := time.Now().Add(-48 * time.Hour)
	modified := time.Now().Add(-24 * time.Hour)
	hash, _ := hashFile(path)
	previous := &Entry{Created: created, Modified: modified, Hash: hash}

	entry, err := buildFileEntry(node, previous, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if !entry.Created.Equal(created) || !entry.Modified.Equal(modified) {
		t.Errorf("expected dates unchanged, got %v / %v", entry.Created, entry.Modified)
	}
}

func TestBuildFileEntry_ChangedContent_KeepsCreatedBumpsModified(t *testing.T) {
	dir := testDir(t)
	path := filepath.Join(dir, "1_x.json")
	writeTestFile(t, path, `{"new":true}`)

	parent, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}
	node := parent.Children[0]
	created := time.Now().Add(-48 * time.Hour)
	previous := &Entry{Created: created, Modified: created, Hash: "some-old-hash"}
	now := time.Now()

	entry, err := buildFileEntry(node, previous, now)
	if err != nil {
		t.Fatal(err)
	}
	if !entry.Created.Equal(created) {
		t.Errorf("expected created to be preserved, got %v", entry.Created)
	}
	if !entry.Modified.Equal(now) {
		t.Errorf("expected modified to bump to now, got %v", entry.Modified)
	}
}

func TestBuildProjectEntry_FolderRollsUpModifiedFromChangedChild(t *testing.T) {
	dir := testDir(t)
	writeTestFile(t, filepath.Join(dir, "test", "1_a.json"), `{}`)
	writeTestFile(t, filepath.Join(dir, "test", "2_b.json"), `{}`)

	parent, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}
	node := parent.Children[0]

	created := time.Now().Add(-48 * time.Hour)
	unchangedTime := time.Now().Add(-24 * time.Hour)
	hashA, _ := hashFile(filepath.Join(dir, "test", "1_a.json"))
	previous := &Entry{
		Created:  created,
		Modified: unchangedTime,
		Contents: map[string]*Entry{
			"1_a.json": {Created: created, Modified: unchangedTime, Hash: hashA},
			"2_b.json": {Created: created, Modified: unchangedTime, Hash: "stale-hash"},
		},
	}
	now := time.Now()

	entry, err := buildProjectEntry(node, previous, now)
	if err != nil {
		t.Fatal(err)
	}
	if !entry.Created.Equal(created) {
		t.Errorf("expected folder created preserved, got %v", entry.Created)
	}
	if !entry.Modified.Equal(now) {
		t.Errorf("expected folder modified to bump since 2_b.json changed, got %v", entry.Modified)
	}
	if !entry.Contents["1_a.json"].Modified.Equal(unchangedTime) {
		t.Errorf("expected 1_a.json's own modified to stay unchanged")
	}
}

func TestBuildProjectCatalogEntry_ReflectsWhichCategoriesExist(t *testing.T) {
	dir := testDir(t)
	writeTestFile(t, filepath.Join(dir, "test", "1_x.json"), `{}`)
	writeTestFile(t, filepath.Join(dir, "tutorial", "1_x.json"), `{}`)

	node, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}

	entry := buildProjectCatalogEntry(node, nil, time.Now())
	if !entry.HasTest || !entry.HasTutorial || entry.HasPreset {
		t.Errorf("expected has_test=true, has_tutorial=true, has_preset=false, got %+v", entry)
	}
}

func TestBuildProjectCatalogEntry_UnchangedFlags_KeepsModified(t *testing.T) {
	dir := testDir(t)
	writeTestFile(t, filepath.Join(dir, "test", "1_x.json"), `{}`)

	node, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}
	modified := time.Now().Add(-24 * time.Hour)
	previous := &Entry{Created: modified, Modified: modified, HasTest: true}

	entry := buildProjectCatalogEntry(node, previous, time.Now())
	if !entry.Modified.Equal(modified) {
		t.Errorf("expected modified unchanged when flags match, got %v", entry.Modified)
	}
}

func TestBuildRootEntry_RollsUpThroughUsernameAndProjectLevels(t *testing.T) {
	dir := testDir(t)
	writeTestFile(t, filepath.Join(dir, "digiconvent", "d9t", "test", "1_x.json"), `{}`)

	root, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	entry := buildRootEntry(root, nil, now)

	if !entry.Modified.Equal(now) {
		t.Fatalf("expected root modified to be now on first build, got %v", entry.Modified)
	}
	project := entry.Contents["digiconvent"].Contents["d9t"]
	if project == nil {
		t.Fatal("expected a digiconvent/d9t entry")
	}
	if !project.HasTest || project.HasTutorial || project.HasPreset {
		t.Errorf("expected has_test=true only, got %+v", project)
	}
	if project.URL != rawURL(filepath.Join(dir, "digiconvent", "d9t")+"/index.json") {
		t.Errorf("unexpected project url: %s", project.URL)
	}
}

func TestBuildRootEntry_NoChanges_KeepsModified(t *testing.T) {
	dir := testDir(t)
	writeTestFile(t, filepath.Join(dir, "digiconvent", "d9t", "test", "1_x.json"), `{}`)

	root, err := walkNode(dir, dir)
	if err != nil {
		t.Fatal(err)
	}

	stableTime := time.Now().Add(-24 * time.Hour)
	previous := buildRootEntry(root, nil, stableTime)

	entry := buildRootEntry(root, previous, time.Now())
	if !entry.Modified.Equal(stableTime) {
		t.Errorf("expected modified to stay stable when nothing changed, got %v", entry.Modified)
	}
}
