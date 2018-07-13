package bundle

import (
	"fmt"
	"strings"

	"github.com/getantibody/antibody/project"
)

type pathBundle struct {
	Project project.Project
}

func (bundle pathBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}

	return fmt.Sprintf(
		`export PATH="%s:$PATH"`,
		strings.Join(bundle.Project.Folders(), ":"),
	), err
}
