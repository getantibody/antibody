package project

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/getantibody/folder"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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
	return gitProject{
		folder: filepath.Join(home, folderName),
		URL:    folder.ToURL(folderName),
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
	return gitProject{
		Version: version,
		URL:     url,
		folder:  filepath.Join(cwd, folder.FromURL(url)),
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
		var w bytes.Buffer
		if _, err := git.PlainClone(g.folder, false, &git.CloneOptions{
			URL:               g.URL,
			ReferenceName:     plumbing.NewBranchReferenceName(g.Version),
			SingleBranch:      true,
			NoCheckout:        false,
			Depth:             1,
			RecurseSubmodules: 1,
			Progress:          &w,
		}); err != nil {
			log.Println("git clone failed for", g.URL, w.String())
			return err
		}
	}
	return nil
}

func (g gitProject) Update() error {
	log.Println("updating:", g.URL)

	repo, err := git.PlainOpen(g.folder)
	if err != nil {
		return err
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	oldRev, err := repo.ResolveRevision(plumbing.Revision(plumbing.HEAD))
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	var w bytes.Buffer
	if err := wt.Pull(&git.PullOptions{
		RemoteName:        "origin",
		ReferenceName:     ref.Name(),
		SingleBranch:      true,
		Depth:             1,
		RecurseSubmodules: 1,
		Progress:          &w,
		Force:             true,
	}); err != git.NoErrAlreadyUpToDate {
		log.Println("git update failed for", g.folder, g.Version, w.String())
		return err
	}

	rev, err := repo.ResolveRevision(plumbing.Revision(plumbing.HEAD))
	if err != nil {
		return err
	}

	if rev.String() != oldRev.String() {
		log.Println("updated:", g.URL, oldRev, "->", rev)
	}

	return nil
}

func (g gitProject) Path() string {
	return filepath.Join(g.folder, g.inner)
}
