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

func (b gitBundle) Download() (string, error) {
	if _, err := os.Stat(b.folder); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--depth", "1", b.url, b.folder)
		return b.folder, clone.Run()
	}
	return b.folder, nil
}

func (b gitBundle) Update() (string, error) {
	pull := exec.Command("git", "-C", b.folder, "pull", "origin", "master")
	return b.folder, pull.Run()
}

type gitBundle struct {
	url, folder string
}

func NewGithubBundle(bundle, home string) Bundle {
	return gitBundle{GH + bundle, folder(bundle, home)}
}
