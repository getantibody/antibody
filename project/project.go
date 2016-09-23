package project

import (
	"io/ioutil"

	"golang.org/x/sync/errgroup"
)

// Project is basically any kind of project (git, local, svn, bzr, nfs...)
type Project interface {
	Download() error
	Update() error
	Folder() string
}

// New decides what kind of project it is, based on the given line
func New(home, line string) Project {
	if line[0] == '/' {
		return NewLocal(line)
	}
	return NewGit(home, line)
}

// List all projects in the given folder
func List(home string) (result []string, err error) {
	entries, err := ioutil.ReadDir(home)
	if err != nil {
		return result, err
	}
	for _, entry := range entries {
		if entry.Mode().IsDir() && entry.Name()[0] != '.' {
			result = append(result, entry.Name())
		}
	}
	return result, nil
}

// Update all projects in the given folder
func Update(home string) error {
	folders, err := List(home)
	if err != nil {
		return err
	}
	var g errgroup.Group
	for _, folder := range folders {
		folder := folder
		g.Go(func() error {
			return NewClonedGit(home, folder).Update()
		})
	}
	return g.Wait()
}
