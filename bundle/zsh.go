package bundle

import (
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

var zshGlobs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

func (bundle zshBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	for _, glob := range zshGlobs {
		files, _ := filepath.Glob(filepath.Join(bundle.Project.Folder(), glob))
		if files == nil {
			continue
		}
		var lines []string
		for _, file := range files {
			lines = append(lines, "source "+file)
		}

		return strings.Join(lines, "\n"), err
	}

	return result, nil
}
