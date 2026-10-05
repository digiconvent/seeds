package tree

import "fmt"

func requireValidName(n *Node) []error {
	if n.IsDir {
		if !folderNamePattern.MatchString(n.Name) {
			return []error{fmt.Errorf("%s: folder name must be alphanumeric/underscore", n.Path)}
		}
		return nil
	}

	name := trimJsonExt(n.Name)
	if name == n.Name {
		return []error{fmt.Errorf("%s: seed files must end in .json", n.Path)}
	}
	if !fileNamePattern.MatchString(name) {
		return []error{fmt.Errorf("%s: file name must match <number>_<alphanumeric_and_underscore>.json", n.Path)}
	}
	return nil
}
