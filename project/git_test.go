package project_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestDownloadAllKinds(t *testing.T) {
	assert := assert.New(t)
	urls := []string{
		"caarlos0/ports",
		"http://github.com/caarlos0/ports",
		"http://github.com/caarlos0/ports.git",
		"https://github.com/caarlos0/ports",
		"https://github.com/caarlos0/ports.git",
		"git://github.com/caarlos0/ports.git",
		"ssh://git@github.com/caarlos0/ports.git",
		"git@github.com:caarlos0/ports.git",
	}
	for _, url := range urls {
		home := home()
		defer os.RemoveAll(home)
		assert.NoError(
			project.NewGit(home, url, "master").Download(),
			"Repo "+url+" failed to download",
		)
	}
}

func TestDownloadAnotherBranch(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(project.NewGit(home, "caarlos0/jvm", "gh-pages").Download())
}

func TestDownloadAndUpdate(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.NewGit(home, "caarlos0/ports", "master")
	assert.NoError(repo.Download())
	assert.NoError(repo.Update())
}

func TestUpdateNonExistentLocalRepo(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.NewGit(home, "caarlos0/ports", "master")
	assert.Error(repo.Update())
}

func TestDownloadNonExistenRepo(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.NewGit(home, "doesn-not-exist-really", "also-nope")
	assert.Error(repo.Download())
}

func TestDownloadMultipleTimes(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.NewGit(home, "caarlos0/ports", "master")
	assert.NoError(repo.Download())
	assert.NoError(repo.Download())
}

func TestDownloadFolderNaming(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.NewGit(home, "caarlos0/ports", "master")
	assert.Equal(
		home+"https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports",
		repo.Folder(),
	)
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
