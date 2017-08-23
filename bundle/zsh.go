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

	base := filepath.Base(bundle.Project.Folder())
	var lines []string

	pre := os.Getenv("ANTIBODY_BUNDLE_PRE")
	if len(pre) > 0 {
		for _, glob := range zshGlobs {
			files, _ := filepath.Glob(filepath.Join(pre, base, glob))
			if files == nil {
				continue
			}
			for _, file := range files {
				lines = append(lines, "source "+file)
			}
			break
		}
	}

	for _, glob := range zshGlobs {
		files, _ := filepath.Glob(filepath.Join(bundle.Project.Folder(), glob))
		if files == nil {
			continue
		}
		for _, file := range files {
			lines = append(lines, "source "+file)
		}
		break
	}

	post := os.Getenv("ANTIBODY_BUNDLE_POST")
	if len(post) > 0 {
		for _, glob := range zshGlobs {
			files, _ := filepath.Glob(filepath.Join(post, base, glob))
			if files == nil {
				continue
			}
			for _, file := range files {
				lines = append(lines, "source "+file)
			}
			break
		}
	}

	result = strings.Join(lines, "\n")
	return result, nil
}
