package tree

func childNames(n *Node) []string {
	names := make([]string, len(n.Children))
	for i, c := range n.Children {
		names[i] = c.Name
	}
	return names
}
