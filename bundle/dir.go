package bundle

import "os"

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
