package tree

import (
	"fmt"
	"slices"
)

func requireKnownCategories(n *Node) []error {
	hasCategory := false
	for _, child := range n.Children {
		if child.IsDir {
			if !slices.Contains(requiredCategories, child.Name) {
				return []error{fmt.Errorf("%s: expected only %v, found %v", n.Path, requiredCategories, childNames(n))}
			}
			hasCategory = true
			continue
		}
		if !slices.Contains(generatedIndexFiles, child.Name) {
			return []error{fmt.Errorf("%s: expected only %v, found %v", n.Path, requiredCategories, childNames(n))}
		}
	}
	if !hasCategory {
		return []error{fmt.Errorf("%s: expected at least one of %v, found none", n.Path, requiredCategories)}
	}
	return nil
}
