package bundle

import (
	"fmt"

	"github.com/getantibody/antibody/project"
)

type fpathBundle struct {
	Project project.Project
}

func (bundle fpathBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	return fmt.Sprintf("fpath+=( %s )", bundle.Project.Path()), err
}
