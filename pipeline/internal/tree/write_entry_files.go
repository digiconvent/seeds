package tree

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func writeEntryFiles(dir string, entry *Entry, render func(*Entry) string) error {
	data, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "index.json"), append(data, '\n'), 0644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(render(entry)), 0644)
}
