package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Repo represents a git repository
type Repo struct {
	url, name, folder string
}

// NewGitRepo creates a new Github Repo with the fullName and local folder.
func NewGitRepo(fullName, folder string) Repo {
	url := urlFor(fullName)
	return Repo{
		name:   fullName,
		url:    url,
		folder: filepath.Join(folder, urlToFolderName(url)),
	}
}

func urlToFolderName(url string) string {
	return strings.Replace(
		strings.Replace(
			url, ":", "-COLON-", -1,
		), "/", "-SLASH-", -1,
	)
}

func urlFor(s string) string {
	var url string
	switch {
	case strings.HasPrefix(s, "http://"):
		fallthrough
	case strings.HasPrefix(s, "https://"):
		fallthrough
	case strings.HasPrefix(s, "git://"):
		fallthrough
	case strings.HasPrefix(s, "ssh://"):
		fallthrough
	case strings.HasPrefix(s, "git@github.com:"):
		url = s
	default:
		url = "https://github.com/" + s
	}
	return url
}

// Folder where the repo was cloned
func (r Repo) Folder() string {
	return r.folder
}

// Name of the repo
func (r Repo) Name() string {
	return r.name
}

// Download clones a repository
func (r Repo) Download() error {
	if _, err := os.Stat(r.folder); os.IsNotExist(err) {
		return exec.Command(
			"git", "clone", "--depth", "1", r.url, r.folder,
		).Run()
	}
	return nil
}

// Update updates a repository
func (r Repo) Update() error {
	return exec.Command("git", "-C", r.folder, "pull", "origin", "master").Run()
}
