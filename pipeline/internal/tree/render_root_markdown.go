package tree

import "strings"

func renderRootMarkdown(entry *Entry) string {
	var b strings.Builder
	b.WriteString("# Seeds\n\n")

	usernames := sortedKeys(entry.Contents)
	for _, username := range usernames {
		b.WriteString("- ")
		b.WriteString(username)
		b.WriteString("\n")

		projects := sortedKeys(entry.Contents[username].Contents)
		for _, project := range projects {
			b.WriteString("  - [")
			b.WriteString(project)
			b.WriteString("](seeds/")
			b.WriteString(username)
			b.WriteString("/")
			b.WriteString(project)
			b.WriteString("/)\n")
		}
	}
	return b.String()
}
