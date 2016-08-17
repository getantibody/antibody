package bundle

import (
	"path/filepath"

	"github.com/getantibody/antibody/event"
	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

var zshGlobs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

func (z zshBundle) Get(events chan event.Event) {
	if err := z.Project.Download(); err != nil {
		events <- event.Error(err)
		return
	}
	for _, glob := range zshGlobs {
		files, err := filepath.Glob(filepath.Join(z.Project.Folder(), glob))
		if err != nil {
			events <- event.Error(err)
			continue
		}
		if files == nil {
			continue
		}
		for _, file := range files {
			events <- event.Shell("source " + file)
			return
		}
	}
}
