package tree

import "time"

func (t *Tree) Generate() *Tree {
	if t.err != nil {
		return t
	}
	now := time.Now()
	for _, usernameNode := range t.Root.Children {
		for _, projectNode := range usernameNode.Children {
			if err := generateProject(projectNode, now); err != nil {
				t.err = err
				return t
			}
		}
	}
	t.err = generateRoot(t.Root, now)
	return t
}
