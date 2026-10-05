package tree

import (
	"encoding/json"
	"fmt"
	"os"
)

func requireValidJSON(n *Node) []error {
	data, err := os.ReadFile(n.Path)
	if err != nil {
		return []error{fmt.Errorf("%s: %w", n.Path, err)}
	}
	if !json.Valid(data) {
		return []error{fmt.Errorf("%s: not valid json", n.Path)}
	}
	return nil
}
