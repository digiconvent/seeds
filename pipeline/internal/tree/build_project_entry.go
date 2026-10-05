package tree

import (
	"slices"
	"time"
)

func buildProjectEntry(node *Node, previous *Entry, now time.Time) (*Entry, error) {
	if !node.IsDir {
		return buildFileEntry(node, previous, now)
	}

	contents := map[string]*Entry{}
	for _, child := range node.Children {
		if slices.Contains(generatedIndexFiles, child.Name) {
			continue
		}
		var prevChild *Entry
		if previous != nil {
			prevChild = previous.Contents[child.Name]
		}
		childEntry, err := buildProjectEntry(child, prevChild, now)
		if err != nil {
			return nil, err
		}
		contents[child.Name] = childEntry
	}
	return foldEntry(contents, previous, now), nil
}
