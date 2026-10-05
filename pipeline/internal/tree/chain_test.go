package tree_test

import (
	"os"
	"testing"

	"github.com/digiconvent/seeds/internal/tree"
)

func TestChain_ValidTree_GeneratesSuccessfully(t *testing.T) {
	chdirToTemp(t)
	validSeedsTree(t)

	if err := tree.Load("seeds").Validate().Generate().Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat("index.json"); err != nil {
		t.Errorf("expected index.json to be written: %v", err)
	}
}

func TestChain_InvalidTree_StopsBeforeGenerating(t *testing.T) {
	chdirToTemp(t)
	writeFile(t, "seeds/digiconvent/d9t/bogus/1_x.json", `{}`)

	err := tree.Load("seeds").Validate().Generate().Err()
	if err == nil {
		t.Fatal("expected the unknown category folder to fail validation")
	}
	if _, statErr := os.Stat("index.json"); statErr == nil {
		t.Error("expected Generate to be skipped after a validation error")
	}
}
