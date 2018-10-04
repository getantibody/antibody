package bundle

import (
	"fmt"
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
	info, err := os.Stat(bundle.Project.Path())
	if err != nil {
		return "", err
	}
	// it is a file, not a folder, so just return it
	if info.Mode().IsRegular() {
		return bundle.Project.Path(), nil
	}
	for _, glob := range zshGlobs {
		files, err := filepath.Glob(filepath.Join(bundle.Project.Path(), glob))
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
		lines = append(lines, fmt.Sprintf("fpath+=( %s )", bundle.Project.Path()))
		return strings.Join(lines, "\n"), err
	}

	return result, nil
}
