package bundle

import (
	"os"
	"path/filepath"

	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

var zshGlobs = []string{"*.zsh", "*.plugin.zsh", "*.sh", "*.zsh-theme"}

func (bundle zshBundle) Get() (result string, err error) {
	if err := bundle.Project.Download(); err != nil {
		return result, err
	}
	for _, glob := range zshGlobs {
		files, _ := filepath.Glob(filepath.Join(bundle.Project.Folder(), glob))
		if files == nil {
			continue
		}
		for _, file := range files {
			stat, _ := os.Lstat(file)
			if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
				continue
			}
			result = "source " + file + ";\n" + result
		}
		return result, nil
	}
	return result, nil
}
