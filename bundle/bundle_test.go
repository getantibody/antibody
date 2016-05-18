package bundle_test

import (
	"os"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestBundleSourceables(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New("caarlos0/zsh-pg", home)
	b.Download()
	assert.NotEmpty(t, b.Sourceables())
}

func TestSourceablesWithoutDownload(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New("caarlos0/zsh-pg", home)
	assert.Empty(t, b.Sourceables())
}

func TestSourceablesDotSh(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New("rupa/z", home)
	b.Download()
	assert.Len(t, b.Sourceables(), 1)
}

func TestSourceablesZshTheme(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New("caiogondim/bullet-train-oh-my-zsh-theme", home)
	b.Download()
	assert.Len(t, b.Sourceables(), 1)
}

func TestListEmptyFolder(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	assert.Empty(t, bundle.List(home))
}

func TestList(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	bundle.New("caarlos0/zsh-pg", home).Download()
	assert.NotEmpty(t, bundle.List(home))
}

func TestParse(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	s := "caarlos0/zsh-pg\ncaarlos0/zsh-open-pr"
	assert.Len(t, bundle.Parse(s, home), 2)
}

func TestParseWithEmptyLines(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	s := "caarlos0/zsh-pg\n\n  \ncaarlos0/zsh-open-pr"
	assert.Len(t, bundle.Parse(s, home), 2)
}

func TestParseWithComment(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	s := "caarlos0/zsh-pg      # this is a comment"
	assert.Len(t, bundle.Parse(s, home), 1)
}
