package tree

func Load(dir string) *Tree {
	root, err := walkNode(dir, dir)
	return &Tree{Root: root, err: err}
}
