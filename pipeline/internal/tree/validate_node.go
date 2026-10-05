package tree

func validateNode(n *Node, depth int) []error {
	var errs []error

	switch depth {
	case depthRoot, depthUsername:
		errs = append(errs, requireAllDirs(n)...)
	case depthRepo:
		errs = append(errs, requireKnownCategories(n)...)
	case depthCategory:
		errs = append(errs, requireUniformChildren(n)...)
	default:
		errs = append(errs, requireUniformChildren(n)...)
		errs = append(errs, requireValidName(n)...)
		if !n.IsDir {
			errs = append(errs, requireValidJSON(n)...)
		}
	}

	for _, child := range n.Children {
		errs = append(errs, validateNode(child, depth+1)...)
	}
	return errs
}
