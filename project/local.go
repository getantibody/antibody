package project

import (
	"os"
	"strings"
	"os/user"
)

// NewLocal Returns a local project, which can be any folder you want to
func NewLocal(folder string) Project {
	return localProject{folder: PrepareFolder(folder)}
}

// PrepareFolder performs path normalization/expansion on local folder path
func PrepareFolder(folder string) string {
	if strings.HasPrefix(folder, "~/") {
		usr, _ := user.Current()
		folder = strings.Replace(folder, "~", usr.HomeDir, 1)
	}
	return strings.Split(folder, " ")[0]
}

type localProject struct {
	folder string
}

func (l localProject) Download() error {
	_, err := os.Stat(l.folder)
	return err
}

func (l localProject) Update() error {
	return l.Download()
}

func (l localProject) Path() string {
	return l.folder
}
