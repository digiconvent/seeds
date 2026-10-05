package tree

import (
	"path/filepath"
	"time"
)

func generateProject(node *Node, now time.Time) error {
	previous := readPreviousEntry(filepath.Join(node.Path, "index.json"))
	entry, err := buildProjectEntry(node, previous, now)
	if err != nil {
		return err
	}
	return writeEntryFiles(node.Path, entry, renderProjectMarkdown)
}
