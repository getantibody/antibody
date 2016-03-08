package git_test

import (
	"os"
	"testing"

	"github.com/akatrevorjay/antibody/git"
	"github.com/akatrevorjay/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestClonesRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("caarlos0/env", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestUpdatesRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("caarlos0/zsh-pg", home)
	assert.NoError(t, repo.Download())
	assert.NoError(t, repo.Update())
	internal.AssertFileCount(t, 1, home)
}

func TestCloneDoesNothingIfFolderAlreadyExists(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("caarlos0/zsh-add-upstream", home)
	assert.NoError(t, repo.Download())
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesUnexistentRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("doesn-not-exist-really", home)
	assert.Error(t, repo.Download())
	internal.AssertFileCount(t, 0, home)
}

func TestUpdatesUnexistentRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("doesn-not-exist-really", home)
	assert.Error(t, repo.Update())
	internal.AssertFileCount(t, 0, home)
}

func TestGetsRepoInfo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGithubRepo("caarlos0/zsh-pg", home)
	assert.Equal(t, "caarlos0/zsh-pg", repo.Name())
	assert.Equal(t, home+"caarlos0-zsh-pg", repo.Folder())
}
