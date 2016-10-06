package bundle

import "github.com/getantibody/antibody/project"

type pathBundle struct {
	Project project.Project
}

func (bundle pathBundle) Get() (result string, err error) {
	if err := bundle.Project.Download(); err != nil {
		return result, err
	}
	return "export PATH=\"" + bundle.Project.Folder() + ":$PATH\"", err
}
