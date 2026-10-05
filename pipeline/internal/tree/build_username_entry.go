package tree

import "time"

func buildUsernameEntry(node *Node, previous *Entry, now time.Time) *Entry {
	contents := map[string]*Entry{}
	for _, projectNode := range node.Children {
		var prevProject *Entry
		if previous != nil {
			prevProject = previous.Contents[projectNode.Name]
		}
		contents[projectNode.Name] = buildProjectCatalogEntry(projectNode, prevProject, now)
	}
	return foldEntry(contents, previous, now)
}
