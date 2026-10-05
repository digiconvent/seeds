package tree

import "regexp"

var folderNamePattern = regexp.MustCompile(`^[A-Za-z0-9_]+$`)
var fileNamePattern = regexp.MustCompile(`^[0-9]+_[A-Za-z0-9_]+$`)
