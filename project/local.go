package project

import (
	"os"
	"strings"
)

// NewLocal Returns a local project, which can be any folder you want to
func NewLocal(line string) (Project, error) {
	folder, err := expandFolder(strings.Split(line, " ")[0])
	return localProject{folder: folder}, err
}

func expandFolder(folder string) (string, error) {
	if strings.HasPrefix(folder, "~/") {
		dir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return strings.Replace(folder, "~", dir, 1), nil
	}
	return folder, nil
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
