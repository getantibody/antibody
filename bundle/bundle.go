package bundle

import (
	"strings"

	"github.com/getantibody/antibody/event"
	"github.com/getantibody/antibody/project"
)

// Bundle main interface.
type Bundle interface {
	Get(events chan event.Event)
}

// New bundle with at the given home (when apply) and line.
//
// Accepted line formats:
//
// - Local bundle (download and update do nothing):
// 		/home/carlos/Code/my-local-bundle
// - Github repo in the owner/repo format:
//		caarlos0/github-repo
// - Git repo in any valid URL form:
//		https://github.com/caarlos0/other-github-repo.git
// - Any type of repo, specifying the kind of resource:
//		caarlos0/add-to-path-style kind:path
// - Any git repo, specifying a branch:
//		caarlos0/versioned-with-branch branch:v1.0 kind:zsh
func New(home, line string) Bundle {
	identifier, kind, version := extract(line)
	proj := projectFor(identifier, version, home)
	if kind == "path" {
		return pathBundle{proj}
	}
	return zshBundle{proj}
}

func projectFor(identifier, version, home string) project.Project {
	if identifier[0] == '/' {
		return project.NewLocal(identifier)
	}
	return project.NewGit(home, identifier, version)
}

func extract(line string) (string, string, string) {
	var version = "master"
	var kind = "zsh"
	parts := strings.Split(line, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, "branch:") {
			version = strings.Replace(part, "branch:", "", -1)
		} else if strings.HasPrefix(part, "kind:") {
			kind = strings.Replace(part, "kind:", "", -1)
		}
	}
	return parts[0], kind, version
}
