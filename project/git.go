package project

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/getantibody/folder"
)

type gitProject struct {
	URL     string
	Version string
	folder  string
	inner   string
}

// NewClonedGit is a git project that was already cloned, so, only Update
// will work here.
func NewClonedGit(home, folderName string) Project {
	version, err := branch(folderName)
	if err != nil {
		version = "master"
	}
	url := folder.ToURL(folderName)
	return gitProject{
		folder:  filepath.Join(home, folderName),
		Version: version,
		URL:     url,
	}
}

// NewGit A git project can be any repository in any given branch. It will
// be downloaded to the provided cwd
func NewGit(cwd, line string) Project {
	version := "master"
	inner := ""
	parts := strings.Split(line, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, "branch:") {
			version = strings.Replace(part, "branch:", "", -1)
		}
		if strings.HasPrefix(part, "folder:") {
			inner = strings.Replace(part, "folder:", "", -1)
		}
	}
	repo := parts[0]
	url := "https://github.com/" + repo
	switch {
	case strings.HasPrefix(repo, "http://"):
		fallthrough
	case strings.HasPrefix(repo, "https://"):
		fallthrough
	case strings.HasPrefix(repo, "git://"):
		fallthrough
	case strings.HasPrefix(repo, "ssh://"):
		fallthrough
	case strings.HasPrefix(repo, "git@gitlab.com:"):
		fallthrough
	case strings.HasPrefix(repo, "git@github.com:"):
		url = repo
	}
	folder := filepath.Join(cwd, folder.FromURL(url))
	return gitProject{
		Version: version,
		URL:     url,
		folder:  folder,
		inner:   inner,
	}
}

func (g gitProject) Download() error {
	if _, err := os.Stat(g.folder); os.IsNotExist(err) {
		// #nosec
		var cmd = exec.Command("git", "clone",
			"--recursive",
			"--depth", "1",
			"-b", g.Version,
			g.URL,
			g.folder)
		cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")
		if bts, err := cmd.CombinedOutput(); err != nil {
			log.Println("git clone failed for", g.URL, string(bts))
			return err
		}
	}
	return nil
}

func (g gitProject) Update() error {
	// #nosec
	if bts, err := exec.Command(
		"git", "-C", g.folder, "pull",
		"--recurse-submodules",
		"origin",
		g.Version,
	).CombinedOutput(); err != nil {
		log.Println("git update failed for", g.folder, string(bts))
		return err
	}
	return nil
}

func branch(folder string) (string, error) {
	// #nosec
	branch, err := exec.Command(
		"git", "-C", folder, "rev-parse", "--abbrev-ref", "HEAD",
	).Output()
	return strings.Replace(string(branch), "\n", "", -1), err
}

func (g gitProject) Folder() string {
	return filepath.Join(g.folder, g.inner)
}
