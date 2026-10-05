package tree

import "errors"

func (t *Tree) Validate() *Tree {
	if t.err != nil {
		return t
	}
	t.err = errors.Join(validateNode(t.Root, depthRoot)...)
	return t
}
