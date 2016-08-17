package project

import "io/ioutil"

// Project is basically any kind of project (git, local, svn, bzr, nfs...)
type Project interface {
	Download() error
	Update() error
	Folder() string
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
	for _, folder := range folders {
		if err := NewClonedGit(home, folder).Update(); err != nil {
			return err
		}
	}
	return nil
}
