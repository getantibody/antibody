package bundle

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

var zshGlobs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

func (bundle zshBundle) Get() (string, error) {
	if err := bundle.Project.Download(); err != nil {
		return "", err
	}
	var lines = []string{}
	for _, folder := range bundle.Project.Folders() {
		for _, glob := range zshGlobs {
			fmt.Println("vaaaaaai", folder)
			files, err := filepath.Glob(filepath.Join(folder, glob))
			if err != nil {
				return "", err
			}
			if files == nil {
				continue
			}
			for _, file := range files {
				lines = append(lines, "source "+file)
			}
			lines = append(lines, fmt.Sprintf("fpath+=( %s )", folder))
			break
		}
	}

	return strings.Join(lines, "\n"), nil
}
