package bundle

import (
	"os"
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
	info, err := os.Stat(bundle.Project.Folder())
	if err != nil {
		return "", err
	}
	// it is a file, not a folder, so just return it
	if info.Mode().IsRegular() {
		return bundle.Project.Folder(), nil
	}
	for _, glob := range zshGlobs {
		files, err := filepath.Glob(filepath.Join(bundle.Project.Folder(), glob))
		if err != nil {
			return result, err
		}
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
