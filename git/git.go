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
	repo := Repo{
		name: fullName,
	}
	switch {
	case strings.HasPrefix(fullName, "http://"):
		fallthrough
	case strings.HasPrefix(fullName, "https://"):
		fallthrough
	case strings.HasPrefix(fullName, "git://"):
		fallthrough
	case strings.HasPrefix(fullName, "ssh://"):
		fallthrough
	case strings.HasPrefix(fullName, "git@github.com:"):
		repo.url = fullName
	default:
		repo.url = "https://github.com/" + fullName
	}
	f := strings.Replace(repo.url, ":", "-COLON-", -1)
	f = strings.Replace(f, "/", "-SLASH-", -1)
	repo.folder = filepath.Join(folder, f)
	return repo
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
