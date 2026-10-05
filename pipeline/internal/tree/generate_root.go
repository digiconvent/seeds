package tree

import "time"

func generateRoot(node *Node, now time.Time) error {
	previous := readPreviousEntry("index.json")
	entry := buildRootEntry(node, previous, now)
	return writeEntryFiles(".", entry, renderRootMarkdown)
}
