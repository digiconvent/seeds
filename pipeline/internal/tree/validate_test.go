package tree_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/digiconvent/seeds/internal/tree"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func chdirToTemp(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(orig) })
}

func validSeedsTree(t *testing.T) {
	t.Helper()
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/1_quorum_not_met.json", `{"import":[]}`)
	writeFile(t, "seeds/digiconvent/d9t/tutorial/iam/1_create_a_user.json", `{"import":[]}`)
	writeFile(t, "seeds/digiconvent/d9t/preset/organisation/corporate/1_board.json", `{"import":[]}`)
}

func TestValidate_WellFormedTree_HasNoErrors(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)

	if err := tree.Load("seeds").Validate().Err(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidate_MixedFilesAndFolders_IsRejected(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)
	writeFile(t, "seeds/digiconvent/d9t/preset/organisation/1_stray_file.json", `{}`)

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "both folders and files") {
		t.Fatalf("expected a mixed-folder error, got %v", err)
	}
}

func TestValidate_GeneratedIndexFilesAlongsideCategories_IsAccepted(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)
	writeFile(t, "seeds/digiconvent/d9t/index.json", `{}`)
	writeFile(t, "seeds/digiconvent/d9t/README.md", `# Seeds`)

	if err := tree.Load("seeds").Validate().Err(); err != nil {
		t.Fatalf("expected generated index files to be accepted, got %v", err)
	}
}

func TestValidate_SubsetOfCategories_IsAccepted(t *testing.T) {
	chdirToTemp(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/1_x.json", `{}`)
	writeFile(t, "seeds/digiconvent/d9t/tutorial/iam/1_x.json", `{}`)

	if err := tree.Load("seeds").Validate().Err(); err != nil {
		t.Fatalf("expected a subset of categories to be valid, got %v", err)
	}
}

func TestValidate_NoCategories_IsRejected(t *testing.T) {
	chdirToTemp(t)
	if err := os.MkdirAll("seeds/digiconvent/d9t", 0755); err != nil {
		t.Fatal(err)
	}

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "expected at least one") {
		t.Fatalf("expected a no-categories error, got %v", err)
	}
}

func TestValidate_UnknownCategoryName_IsRejected(t *testing.T) {
	chdirToTemp(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/1_x.json", `{}`)
	writeFile(t, "seeds/digiconvent/d9t/bogus/1_x.json", `{}`)

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "expected only") {
		t.Fatalf("expected an unknown-category error, got %v", err)
	}
}

func TestValidate_BadFileName_IsRejected(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/not_numbered.json", `{}`)

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "must match") {
		t.Fatalf("expected a naming error, got %v", err)
	}
}

func TestValidate_InvalidJSON_IsRejected(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/2_broken.json", `{not json`)

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "not valid json") {
		t.Fatalf("expected a JSON error, got %v", err)
	}
}

func TestValidate_FileWhereRepoFolderExpected_IsRejected(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)
	writeFile(t, "seeds/digiconvent/stray.json", `{}`)

	err := tree.Load("seeds").Validate().Err()
	if !containsSubstring(err, "expected a directory") {
		t.Fatalf("expected a directory-expected error, got %v", err)
	}
}

func containsSubstring(err error, substr string) bool {
	return err != nil && strings.Contains(err.Error(), substr)
}
