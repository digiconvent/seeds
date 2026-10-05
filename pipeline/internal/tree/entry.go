package tree

import "time"

type Entry struct {
	Created     time.Time         `json:"created"`
	Modified    time.Time         `json:"modified"`
	URL         string            `json:"url,omitempty"`
	Hash        string            `json:"hash,omitempty"`
	Contents    map[string]*Entry `json:"contents,omitempty"`
	HasTest     bool              `json:"has_test,omitempty"`
	HasTutorial bool              `json:"has_tutorial,omitempty"`
	HasPreset   bool              `json:"has_preset,omitempty"`
}
