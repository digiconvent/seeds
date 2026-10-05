package tree

import "path"

func relPath(root, dir string) string {
	if dir == root {
		return path.Base(root)
	}
	rel := dir[len(root)+1:]
	return path.Join(path.Base(root), rel)
}
