package tree

import "time"

func buildFileEntry(node *Node, previous *Entry, now time.Time) (*Entry, error) {
	hash, err := hashFile(node.Path)
	if err != nil {
		return nil, err
	}

	entry := &Entry{URL: rawURL(node.Path), Hash: hash}
	if previous != nil {
		entry.Created = previous.Created
	} else {
		entry.Created = now
	}
	if previous != nil && previous.Hash == hash {
		entry.Modified = previous.Modified
	} else {
		entry.Modified = now
	}
	return entry, nil
}
