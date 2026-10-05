package tree

import "fmt"

func requireAllDirs(n *Node) []error {
	var errs []error
	for _, child := range n.Children {
		if !child.IsDir {
			errs = append(errs, fmt.Errorf("%s: expected a directory but found a file", child.Path))
		}
	}
	return errs
}
