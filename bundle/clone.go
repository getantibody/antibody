package bundle

import "github.com/getantibody/antibody/project"

type cloneBundle struct {
	Project project.Project
}

func (bundle cloneBundle) Get() (result string, err error) {
	err = bundle.Project.Download()
	return result, err
}
