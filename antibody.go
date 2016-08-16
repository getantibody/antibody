package antibody

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/caarlos0/gohome"
)

type Antibody struct {
	Events  chan Event
	Bundles []string
	Home    string
}

func New(bundles []string) *Antibody {
	return &Antibody{
		Bundles: bundles,
		Events:  make(chan Event, len(bundles)),
	}
}

type Event struct {
	Shell string
	Error error
}

func (a *Antibody) Bundle() (sh string, err error) {
	var shs []string
	var total = len(a.Bundles)
	var count int
	done := make(chan bool)

	for _, sh := range a.Bundles {
		go func(s string) {
			NewBundle(s).Get(a.Events)
			done <- true
		}(sh)
	}
	for {
		select {
		case event := <-a.Events:
			if event.Error != nil {
				return "", err
			}
			shs = append(shs, event.Shell)
		case <-done:
			count++
			if count == total {
				return strings.Join(shs, "\n"), nil
			}
		}
	}
}

// Bundle interface.
//
// Accepted formats:
//
// - Local bundle (download and update do nothing):
// 		/home/carlos/Code/my-local-bundle
// - Github repo in the owner/repo format:
//		caarlos0/github-repo
// - Git repo in any valid URL form:
//		https://github.com/caarlos0/other-github-repo.git
// - Any type of repo, specifying the kind of resource:
//		caarlos0/add-to-path-style kind:Path
// - Any git repo, specifying a branch:
//		caarlos0/versioned-with-branch branch:v1.0
type Bundle interface {
	Get(events chan Event)
}

type Project interface {
	Download() error
	Update() error
	Folder() string
}

func NewBundle(sh string) Bundle {
	var project Project
	var bundle Bundle
	var version = "master"
	var kind = "zsh"
	parts := strings.Split(sh, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, "branch:") {
			version = strings.Replace(part, "branch:", "", -1)
		} else if strings.HasPrefix(part, "kind:") {
			kind = strings.Replace(part, "kind:", "", -1)
		}
	}
	if sh[0] == '/' {
		project = LocalProject{
			folder: parts[0],
		}
	} else {
		project = NewGitProject(parts[0], version)
	}
	if kind == "zsh" {
		bundle = ZshBundle{project}
	} else if kind == "path" {
		bundle = PathBundle{project}
	}
	return bundle
}

type LocalProject struct {
	folder string
}

func (l LocalProject) Download() error {
	return nil
}

func (l LocalProject) Update() error {
	return nil
}

func (l LocalProject) Folder() string {
	return l.folder
}

type GitProject struct {
	URL     string
	Version string
	folder  string
}

func NewGitProject(repo, version string) GitProject {
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
	folder := gohome.Cache("__antibody") + "/" + strings.Replace(
		strings.Replace(
			url, ":", "-COLON-", -1,
		), "/", "-SLASH-", -1,
	)
	log.Println(folder)
	return GitProject{
		Version: version,
		URL:     url,
		folder:  folder,
	}
}

func (g GitProject) Download() error {
	if _, err := os.Stat(g.folder); os.IsNotExist(err) {
		return exec.Command(
			"git", "clone", "--depth", "1", "-b", g.Version, g.URL, g.folder,
		).Run()
	}
	return nil
}

func (g GitProject) Update() error {
	return exec.Command(
		"git", "-C", g.folder, "pull", "origin", g.Version,
	).Run()
}

func (g GitProject) Folder() string {
	return g.folder
}

type ZshBundle struct {
	Project Project
}

var zshGlobs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

func (z ZshBundle) Get(events chan Event) {
	if err := z.Project.Download(); err != nil {
		events <- Event{Error: err}
		return
	}
	for _, glob := range zshGlobs {
		files, err := filepath.Glob(filepath.Join(z.Project.Folder(), glob))
		if err != nil {
			events <- Event{Error: err}
			continue
		}
		if files == nil {
			continue
		}
		for _, file := range files {
			events <- Event{Shell: "source " + file}
			return
		}
	}
}

type PathBundle struct {
	Project Project
}

func (z PathBundle) Get(events chan Event) {
	if err := z.Project.Download(); err != nil {
		events <- Event{Error: err}
		return
	}
	events <- Event{Shell: "export PATH=\"" + z.Project.Folder() + ":$PATH\""}
}
