package tree

import "time"

func foldEntry(contents map[string]*Entry, previous *Entry, now time.Time) *Entry {
	changed := previous == nil
	for _, e := range contents {
		if e.Modified.Equal(now) {
			changed = true
		}
	}

	entry := &Entry{Contents: contents}
	if previous != nil {
		entry.Created = previous.Created
	} else {
		entry.Created = now
	}
	if changed {
		entry.Modified = now
	} else {
		entry.Modified = previous.Modified
	}
	return entry
}
