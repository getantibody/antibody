package antibody

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (b gitBundle) Download() error {
	if _, err := os.Stat(b.folder); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--depth", "1", b.url, b.folder)
		return clone.Run()
	}
	return nil
}

func (b gitBundle) Update() error {
	pull := exec.Command("git", "-C", b.folder, "pull", "origin", "master")
	return pull.Run()
}

func (b gitBundle) Folder() string {
	return b.folder
}

func (b gitBundle) Sourceables() []string {
	globs := []string{"*.plugin.zsh", "*.zsh", "*.sh"}
	for _, glob := range globs {
		files, _ := filepath.Glob(filepath.Join(b.Folder(), glob))
		if files != nil {
			return files
		}
	}
	return nil
}

type gitBundle struct {
	url, folder string
}

// NewGitBundle creates a new Bundle using Github as its source.
func NewGitBundle(bundle, home string) Bundle {
	return gitBundle{
		url:    "https://github.com/" + bundle,
		folder: filepath.Join(home, strings.Replace(bundle, "/", "-", -1)),
	}
}
