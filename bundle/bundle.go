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
func New(home, line string) (Bundle, error) {
	proj, err := project.New(home, line)
	if err != nil {
		return nil, err
	}
	switch kind(line) {
	case "path":
		return pathBundle{Project: proj}, nil
	case "fpath":
		return fpathBundle{Project: proj}, nil
	case "dummy":
		return dummyBundle{Project: proj}, nil
	default:
		return zshBundle{Project: proj}, nil
	}
}

func kind(line string) string {
	for _, part := range strings.Split(line, " ") {
		if strings.HasPrefix(part, "kind:") {
			return strings.Replace(part, "kind:", "", -1)
		}
	}
	return "zsh"
}
