package bundle

import (
	"github.com/getantibody/antibody/helper"
	"github.com/getantibody/antibody/project"
)

type pathBundle struct {
	Project project.Project
}

func (bundle pathBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	return helper.ComposeEnvPath(bundle.Project.Path()), err
}
