package tree

import (
	"encoding/json"
	"os"
)

func readPreviousEntry(path string) *Entry {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var entry Entry
	if json.Unmarshal(data, &entry) != nil {
		return nil
	}
	return &entry
}
