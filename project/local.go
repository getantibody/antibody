package project

import (
	"os"
	"strings"
)

// NewLocal Returns a local project, which can be any folder you want to
func NewLocal(folder string) Project {
	return localProject{folder: strings.Split(folder, " ")[0]}
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

func (l localProject) Folder() string {
	return l.folder
}
