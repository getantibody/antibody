package project

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type gitProject struct {
	URL     string
	Version string
	folder  string
}

// NewGit A git project can be any repository in any given branch. It will
// be downloaded to the provided cwd
func NewGit(cwd, repo, version string) Project {
	var url string
	switch {
	case strings.HasPrefix(repo, "http://"):
		fallthrough
	case strings.HasPrefix(repo, "https://"):
		fallthrough
	case strings.HasPrefix(repo, "git://"):
		fallthrough
	case strings.HasPrefix(repo, "ssh://"):
		fallthrough
	case strings.HasPrefix(repo, "git@github.com:"):
		url = repo
	default:
		url = "https://github.com/" + repo
	}
	folder := filepath.Join(
		cwd,
		strings.Replace(
			strings.Replace(
				url, ":", "-COLON-", -1,
			), "/", "-SLASH-", -1,
		),
	)
	return gitProject{
		Version: version,
		URL:     url,
		folder:  folder,
	}
}

func (g gitProject) Download() error {
	if _, err := os.Stat(g.folder); os.IsNotExist(err) {
		cmd := exec.Command(
			"git", "clone", "--depth", "1", "-b", g.Version, g.URL, g.folder,
		)
		if bts, err := cmd.CombinedOutput(); err != nil {
			log.Println("git clone failed for", g.URL, err, string(bts))
			return err
		}
	}
	return nil
}

func (g gitProject) Update() error {
	return exec.Command(
		"git", "-C", g.folder, "pull", "origin", g.Version,
	).Run()
}

func (g gitProject) Folder() string {
	return g.folder
}
