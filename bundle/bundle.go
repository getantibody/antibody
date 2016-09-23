package bundle

import (
	"strings"

	"github.com/getantibody/antibody/project"
)

// Bundle main interface.
type Bundle interface {
	Get() (result string, err error)
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
	kind := extract(line)
	proj := project.New(home, line)
	if kind == "path" {
		return pathBundle{proj}
	}
	return zshBundle{proj}
}

func extract(line string) string {
	for _, part := range strings.Split(line, " ") {
		if strings.HasPrefix(part, "kind:") {
			return strings.Replace(part, "kind:", "", -1)
		}
	}
	return "zsh"
}
