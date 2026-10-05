package tree

func trimJsonExt(name string) string {
	const ext = ".json"
	if len(name) > len(ext) && name[len(name)-len(ext):] == ext {
		return name[:len(name)-len(ext)]
	}
	return name
}
