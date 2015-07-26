package antibody

import (
	"testing"

	"github.com/caarlos0/antibody/doubles"
	"github.com/stretchr/testify/assert"
)

func TestClonesValidRepo(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	err := bundle.Download()
	expected := home + "caarlos0-zsh-pg"

	assert.Equal(t, expected, bundle.Folder())
	assert.NoError(t, err)
	assertBundledPlugins(t, 1, home)
}

func TestClonesValidRepoTwoTimes(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	err := bundle.Download()
	expected := home + "caarlos0-zsh-pg"
	assert.Equal(t, expected, bundle.Folder())
	assert.NoError(t, err)
	assertBundledPlugins(t, 1, home)
}

func TestClonesInvalidRepo(t *testing.T) {
	home := doubles.TempHome()
	err := NewGitBundle("this-doesnt-exist", home).Download()
	assert.Error(t, err)
}

func TestPullsRepo(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	err := bundle.Update()
	assert.NoError(t, err)
}

func TestSourceablesDotPluginZsh(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	srcs := bundle.Sourceables()
	assert.Len(t, srcs, 1)
}

func TestSourceablesDotSh(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("rupa/z", home)
	bundle.Download()
	srcs := bundle.Sourceables()
	assert.Len(t, srcs, 1)
}
