package antibody

import (
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/project"
)

type Antibody struct {
	Events  chan Event
	Bundles []string
	Home    string
}

func New(home string, bundles []string) *Antibody {
	return &Antibody{
		Bundles: bundles,
		Events:  make(chan Event, len(bundles)),
		Home:    home,
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
			NewBundle(a.Home, s).Get(a.Events)
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

func NewBundle(home, sh string) Bundle {
	var proj project.Project
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
		proj = project.NewLocal(parts[0])
	} else {
		proj = project.NewGit(home, parts[0], version)
	}
	if kind == "zsh" {
		bundle = ZshBundle{proj}
	} else if kind == "path" {
		bundle = PathBundle{proj}
	}
	return bundle
}

type ZshBundle struct {
	Project project.Project
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
	Project project.Project
}

func (z PathBundle) Get(events chan Event) {
	if err := z.Project.Download(); err != nil {
		events <- Event{Error: err}
		return
	}
	events <- Event{Shell: "export PATH=\"" + z.Project.Folder() + ":$PATH\""}
}
