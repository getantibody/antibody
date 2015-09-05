package bundle

import (
	"io/ioutil"
	"path/filepath"

	"github.com/caarlos0/antibody/git"
)

// Bundle is a git-based bundle/plugin
type Bundle struct {
	git.Repo
}

// New creates a new bundle instance
func New(fullName, folder string) Bundle {
	return Bundle{
		git.NewGithubRepo(fullName, folder),
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

// List all bundles in the given folder
func List(folder string) []Bundle {
	entries, _ := ioutil.ReadDir(folder)
	var bundles []Bundle
	for _, bundle := range entries {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			bundles = append(bundles, New(bundle.Name(), folder))
		}
	}
	return bundles
}
