package tree

import "time"

func buildRootEntry(node *Node, previous *Entry, now time.Time) *Entry {
	contents := map[string]*Entry{}
	for _, usernameNode := range node.Children {
		var prevUsername *Entry
		if previous != nil {
			prevUsername = previous.Contents[usernameNode.Name]
		}
		contents[usernameNode.Name] = buildUsernameEntry(usernameNode, prevUsername, now)
	}
	return foldEntry(contents, previous, now)
}
