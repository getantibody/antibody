package bundle

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/git"
)

var globs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

// Bundle is a git-based bundle/plugin
type Bundle struct {
	git.Repo
}

// New creates a new bundle instance
func New(fullName, folder string) Bundle {
	return Bundle{git.NewGithubRepo(fullName, folder)}
}

// Sourceables returns the list of files that could be sourced
func (b Bundle) Sourceables() []string {
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

// Parse a list of bundles, one per line, into a Bundle slice
func Parse(s, folder string) []Bundle {
	var bundles []Bundle
	for _, decl := range strings.Split(s, "\n") {
		b := strings.Split(decl, "#")[0]
		b = strings.TrimSpace(b)
		if b != "" {
			bundles = append(bundles, New(b, folder))
		}
	}
	return bundles
}
