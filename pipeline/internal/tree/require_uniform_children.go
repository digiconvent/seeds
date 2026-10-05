package tree

import "fmt"

func requireUniformChildren(n *Node) []error {
	if !n.IsDir || len(n.Children) == 0 {
		return nil
	}
	dirs, files := 0, 0
	for _, c := range n.Children {
		if c.IsDir {
			dirs++
		} else {
			files++
		}
	}
	if dirs > 0 && files > 0 {
		return []error{fmt.Errorf("%s: contains both folders and files", n.Path)}
	}
	return nil
}
