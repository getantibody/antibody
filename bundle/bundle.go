package bundle

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/git"
)

var globs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

// BundleType is an interface for different types of bundles
type Type interface {
	Folder() string
	Name() string
	Download() error
	Update() error
}

// Bundle is a plugin to install
type Bundle struct {
	Type
}

// DirBundle is a bundle for local plugins
type DirBundle struct {
	name, folder string
}

// Folder where the local bundle exists
func (d DirBundle) Folder() string {
	return d.folder
}

// Name of the local bundle
func (d DirBundle) Name() string {
	return d.name
}

// Download simply checks the local bundle exists
func (d DirBundle) Download() error {
	_, err := os.Stat(d.folder)
	return err
}

// Update is a no-op
func (d DirBundle) Update() error {
	return nil
}

// New creates a new bundle instance
func New(fullName, folder string) Bundle {
	if strings.HasPrefix(fullName, "/") {
		return Bundle{DirBundle{path.Base(fullName), fullName}}
	}
	return Bundle{git.NewGitRepo(fullName, folder)}
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
