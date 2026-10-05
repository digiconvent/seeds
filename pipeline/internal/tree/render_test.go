package tree_test

import (
	"os"
	"strings"
	"testing"

	"github.com/digiconvent/seeds/internal/tree"
)

func TestGenerate_WritesRootAndProjectIndexAndReadme(t *testing.T) {
	chdirToTemp(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/1_quorum_not_met.json", `{}`)

	if err := tree.Load("seeds").Validate().Generate().Err(); err != nil {
		t.Fatal(err)
	}

	rootIndex, err := os.ReadFile("index.json")
	if err != nil {
		t.Fatalf("expected root index.json to be written: %v", err)
	}
	if !strings.Contains(string(rootIndex), `"has_test": true`) {
		t.Errorf("expected root index.json to record has_test, got:\n%s", rootIndex)
	}

	rootReadme, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("expected root README.md to be written: %v", err)
	}
	if !strings.Contains(string(rootReadme), "[d9t](seeds/digiconvent/d9t/)") {
		t.Errorf("expected root README.md to link to the project folder, got:\n%s", rootReadme)
	}

	projectIndex, err := os.ReadFile("seeds/digiconvent/d9t/index.json")
	if err != nil {
		t.Fatalf("expected project index.json to be written: %v", err)
	}
	if !strings.Contains(string(projectIndex), "1_quorum_not_met.json") {
		t.Errorf("expected project index.json to reference the seed file, got:\n%s", projectIndex)
	}

	projectReadme, err := os.ReadFile("seeds/digiconvent/d9t/README.md")
	if err != nil {
		t.Fatalf("expected project README.md to be written: %v", err)
	}
	if !strings.Contains(string(projectReadme), "- 1_quorum_not_met\n") {
		t.Errorf("expected project README.md to list the file without extension, got:\n%s", projectReadme)
	}
	if strings.Contains(string(projectReadme), "1_quorum_not_met.json") {
		t.Errorf("expected the .json extension to be trimmed in the README, got:\n%s", projectReadme)
	}
}

func TestGenerate_RunningTwiceWithNoChanges_LeavesDatesStable(t *testing.T) {
	chdirToTemp(t)
	writeFile(t, "seeds/digiconvent/d9t/test/dtap/1_quorum_not_met.json", `{}`)

	if err := tree.Load("seeds").Validate().Generate().Err(); err != nil {
		t.Fatal(err)
	}
	firstRun, err := os.ReadFile("seeds/digiconvent/d9t/index.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := tree.Load("seeds").Validate().Generate().Err(); err != nil {
		t.Fatal(err)
	}
	secondRun, err := os.ReadFile("seeds/digiconvent/d9t/index.json")
	if err != nil {
		t.Fatal(err)
	}

	if string(firstRun) != string(secondRun) {
		t.Errorf("expected a second run with no changes to produce identical output.\nfirst:\n%s\nsecond:\n%s", firstRun, secondRun)
	}
}
