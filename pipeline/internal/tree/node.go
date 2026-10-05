package tree

type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	IsDir    bool    `json:"is_dir"`
	Children []*Node `json:"children,omitempty"`
}
