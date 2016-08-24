package project

import (
	"io/ioutil"

	"github.com/getantibody/antibody/event"
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
func List(home string) ([]string, error) {
	var result []string
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
	total := len(folders)
	var count int
	events := make(chan event.Event)
	for _, folder := range folders {
		go func(folder string) {
			if err := NewClonedGit(home, folder).Update(); err != nil {
				events <- event.Error(err)
			}
			events <- event.Shell("")
		}(folder)
	}
	for {
		evt := <-events
		if evt.Error != nil {
			return evt.Error
		}
		count++
		if count == total {
			return nil
		}
	}
}
