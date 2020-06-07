package bundle

import (
	"github.com/getantibody/antibody/helper"
	"github.com/getantibody/antibody/project"
)

type fpathBundle struct {
	Project project.Project
}

func (bundle fpathBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	return helper.ComposeFPath(bundle.Project.Path()), err
}
