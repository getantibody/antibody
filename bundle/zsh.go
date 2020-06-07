package bundle

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/helper"
	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

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
		// XXX: should we add the parent folder to fpath too?
		return helper.ComposeSource(bundle.Project.Path()), nil
	}
	for _, glob := range []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"} {
		files, err := filepath.Glob(filepath.Join(bundle.Project.Path(), glob))
		if err != nil {
			return result, err
		}
		if files == nil {
			continue
		}
		var lines []string
		for _, file := range files {
			lines = append(lines, helper.ComposeSource(file))
		}
		lines = append(lines, helper.ComposeFPath(bundle.Project.Path()))
		return strings.Join(lines, "\n"), err
	}

	return result, nil
}
