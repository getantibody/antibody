package antibody

import (
	"path/filepath"

	"github.com/caarlos0/antibody/git"
)

// Bundle is a git-based bundle/plugin
type Bundle struct {
	git.Repo
}

// NewBundle creates a new bundle instance
func NewBundle(bundle, home string) Bundle {
	return Bundle{
		git.NewGithubRepo(bundle, home),
	}
}

// Sourceables returns the list of files that could be sourced
func (b Bundle) Sourceables() []string {
	globs := []string{"*.plugin.zsh", "*.zsh", "*.sh"}
	for _, glob := range globs {
		files, _ := filepath.Glob(filepath.Join(b.Folder(), glob))
		if files != nil {
			return files
		}
	}
	return nil
}
