package bundle

import (
	"github.com/getantibody/antibody/event"
	"github.com/getantibody/antibody/project"
)

type pathBundle struct {
	Project project.Project
}

func (z pathBundle) Get(events chan event.Event) {
	if err := z.Project.Download(); err != nil {
		events <- event.Error(err)
		return
	}
	events <- event.Shell("export PATH=\"" + z.Project.Folder() + ":$PATH\"")
}
