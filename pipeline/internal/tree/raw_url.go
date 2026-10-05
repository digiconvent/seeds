package tree

const rawBaseURL = "https://raw.githubusercontent.com/digiconvent/seeds/main"

func rawURL(path string) string {
	return rawBaseURL + "/" + path
}
