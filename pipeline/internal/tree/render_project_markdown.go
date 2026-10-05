package tree

import "strings"

func renderProjectMarkdown(entry *Entry) string {
	var b strings.Builder
	b.WriteString("# Seeds\n\n")
	renderEntryContents(&b, entry, 0)
	return b.String()
}
