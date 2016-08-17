package naming

import "strings"

var replaces = []struct{ a, b string }{
	{":", "-COLON-"},
	{"/", "-SLASH-"},
	{"@", "-AT-"},
}

// URLToFolder converts the given URL to a folder name
func URLToFolder(url string) string {
	result := url
	for _, replace := range replaces {
		result = strings.Replace(result, replace.a, replace.b, -1)
	}
	return result
}

// FolderToURL converts the given folder to an URL
func FolderToURL(folder string) string {
	result := folder
	for _, replace := range replaces {
		result = strings.Replace(result, replace.b, replace.a, -1)
	}
	return result
}
