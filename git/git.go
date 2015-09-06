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

// NewGithubRepo creates a new Github Repo with the fullName and local folder.
func NewGithubRepo(fullName, folder string) Repo {
	return Repo{
		url:    "https://github.com/" + fullName,
		name:   fullName,
		folder: filepath.Join(folder, strings.Replace(fullName, "/", "-", -1)),
	}
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
