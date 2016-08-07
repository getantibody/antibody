package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Repo represents a git repository
type Repo struct {
	url, name, folder, branch string
}

// NewGitRepo creates a new Github Repo with the line that defines it and
// the local antibody folder.
func NewGitRepo(line, folder string) Repo {
	parts := strings.Split(line, " ")
	branch := "master"
	if len(parts) > 1 {
		branch = parts[1]
	}
	url := urlFor(folderNameToURL(parts[0]))
	return Repo{
		name:   repoNameFor(url),
		url:    url,
		branch: branch,
		folder: filepath.Join(folder, urlToFolderName(url)),
	}
}

func folderNameToURL(url string) string {
	return strings.Replace(
		strings.Replace(
			url, "-COLON-", ":", -1,
		), "-SLASH-", "/", -1,
	)
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

func repoNameFor(s string) string {
	ss := strings.Split(s, "/")
	size := len(ss)
	return ss[size-2] + "/" + ss[size-1]
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
			"git", "clone", "--depth", "1", "-b", r.branch, r.url, r.folder,
		).Run()
	}
	return nil
}

// Update updates a repository
func (r Repo) Update() error {
	branch, err := r.Branch()
	if err != nil {
		return err
	}
	return exec.Command("git", "-C", r.folder, "pull", "origin", branch).Run()
}

// Branch gets the current repo branch name
func (r Repo) Branch() (string, error) {
	branch, err := exec.Command(
		"git", "-C", r.folder, "rev-parse", "--abbrev-ref", "HEAD",
	).Output()
	return strings.Replace(string(branch), "\n", "", -1), err
}
