package tree

import "time"

func buildProjectCatalogEntry(node *Node, previous *Entry, now time.Time) *Entry {
	hasTest, hasTutorial, hasPreset := false, false, false
	for _, c := range node.Children {
		switch c.Name {
		case "test":
			hasTest = true
		case "tutorial":
			hasTutorial = true
		case "preset":
			hasPreset = true
		}
	}

	entry := &Entry{
		URL:         rawURL(node.Path + "/index.json"),
		HasTest:     hasTest,
		HasTutorial: hasTutorial,
		HasPreset:   hasPreset,
	}
	changed := previous == nil ||
		previous.HasTest != hasTest ||
		previous.HasTutorial != hasTutorial ||
		previous.HasPreset != hasPreset
	if previous != nil {
		entry.Created = previous.Created
	} else {
		entry.Created = now
	}
	if changed {
		entry.Modified = now
	} else {
		entry.Modified = previous.Modified
	}
	return entry
}
