package antibody

import (
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func folder(bundle string, home string) string {
	return home + strings.Replace(bundle, "/", "-", -1)
}

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

type gitBundle struct {
	url, folder string
}

func NewGitBundle(bundle, home string) Bundle {
	return gitBundle{GH + bundle, folder(bundle, home)}
}
