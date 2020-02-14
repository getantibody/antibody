package project

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/getantibody/folder"
)

var gitCmdEnv = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0", "GIT_ASKPASS=0")

type gitProject struct {
	URL     string
	Version string
	folder  string
	inner   string
}

// NewClonedGit is a git project that was already cloned, so, only Update
// will work here.
func NewClonedGit(home, folderName string) Project {
	folderPath := filepath.Join(home, folderName)
	version, err := branch(folderPath)
	if err != nil {
		version = "master"
	}
	url := folder.ToURL(folderName)
	return gitProject{
		folder:  folderPath,
		Version: version,
		URL:     url,
	}
}

const (
	branchMarker = "branch:"
	pathMarker   = "path:"
)

// NewGit A git project can be any repository in any given branch. It will
// be downloaded to the provided cwd
func NewGit(cwd, line string) Project {
	version := "master"
	inner := ""
	parts := strings.Split(line, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, branchMarker) {
			version = strings.Replace(part, branchMarker, "", -1)
		}
		if strings.HasPrefix(part, pathMarker) {
			inner = strings.Replace(part, pathMarker, "", -1)
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

// nolint: gochecknoglobals
var locks sync.Map

func (g gitProject) Download() error {
	l, _ := locks.LoadOrStore(g.folder, &sync.Mutex{})
	lock := l.(*sync.Mutex)
	lock.Lock()
	defer lock.Unlock()
	if _, err := os.Stat(g.folder); os.IsNotExist(err) {
		// #nosec
		var cmd = exec.Command("git", "clone",
			"--recursive",
			"--depth", "1",
			"-b", g.Version,
			g.URL,
			g.folder,
		)
		cmd.Env = gitCmdEnv

		if bts, err := cmd.CombinedOutput(); err != nil {
			log.Println("git clone failed for", g.URL, string(bts))
			return err
		}
	}
	return nil
}

func (g gitProject) Update() error {
	fmt.Println("updating:", g.URL)
	// #nosec
	cmd := exec.Command(
		"git", "pull",
		"--recurse-submodules",
		"origin",
		g.Version,
	)
	cmd.Env = gitCmdEnv

	cmd.Dir = g.folder
	if bts, err := cmd.CombinedOutput(); err != nil {
		log.Println("git update failed for", g.folder, string(bts))
		return err
	}
	return nil
}

func branch(folder string) (string, error) {
	// #nosec
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = folder
	branch, err := cmd.Output()
	return strings.Replace(string(branch), "\n", "", -1), err
}

func (g gitProject) Path() string {
	return filepath.Join(g.folder, g.inner)
}
