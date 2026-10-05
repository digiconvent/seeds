package tree

import "strings"

func renderEntryContents(b *strings.Builder, entry *Entry, depth int) {
	for _, name := range sortedKeys(entry.Contents) {
		child := entry.Contents[name]
		b.WriteString(strings.Repeat("  ", depth))
		b.WriteString("- ")
		b.WriteString(trimJsonExt(name))
		b.WriteString("\n")
		renderEntryContents(b, child, depth+1)
	}
}
