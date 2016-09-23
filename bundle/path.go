package bundle

import "github.com/getantibody/antibody/project"

type pathBundle struct {
	Project project.Project
}

func (z pathBundle) Get() (result string, err error) {
	if err := z.Project.Download(); err != nil {
		return result, err
	}
	return "export PATH=\"" + z.Project.Folder() + ":$PATH\"", err
}
