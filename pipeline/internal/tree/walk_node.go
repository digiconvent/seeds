package tree

import (
	"os"
	"path"
	"sort"
)

func walkNode(root, dir string) (*Node, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	node := &Node{Name: path.Base(dir), Path: relPath(root, dir), IsDir: info.IsDir()}
	if !info.IsDir() {
		return node, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	for _, e := range entries {
		child, err := walkNode(root, path.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		node.Children = append(node.Children, child)
	}
	return node, nil
}
