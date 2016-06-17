package git_test

import (
	"os"
	"testing"

	"github.com/getantibody/antibody/git"
	"github.com/getantibody/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestClonesUserSlashRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("caarlos0/jvm", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesHttp(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("http://github.com/caarlos0/jvm", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesHttps(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("https://github.com/caarlos0/jvm", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesGitProtocol(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("git://github.com/caarlos0/jvm.git", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesSshProtocol(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("ssh://github.com/caarlos0/jvm.git", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesSsh(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("git@github.com:caarlos0/jvm.git", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestUpdateAlreadyClonedURL(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	git.NewGitRepo("https://github.com/caarlos0/jvm", home).Download()
	repo := git.NewGitRepo("https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm", home)
	assert.NoError(t, repo.Update())
	internal.AssertFileCount(t, 1, home)
}
