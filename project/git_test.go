package project_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
    "strings"
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
		"https://gitlab.com/caarlos0/test.git",
		// FIXME: those fail on travis:
		// "git@gitlab.com:caarlos0/test.git",
		// "ssh://git@github.com/caarlos0/ports.git",
		// "git@github.com:caarlos0/ports.git",
	}
	for _, url := range urls {
		home := home()
		assert.NoError(
			project.NewGit(home, url).Download(),
			"Repo "+url+" failed to download",
		)
	}
}

func TestDownloadSubmodules(t *testing.T) {
	var assert = assert.New(t)
	var home = home()
	var proj = project.NewGit(home, "fribmendes/geometry")
	var module = filepath.Join(proj.Folder(), "lib/zsh-async")
	assert.NoError(proj.Download())
	assert.NoError(proj.Update())
	files, err := ioutil.ReadDir(module)
	assert.NoError(err)
	assert.True(len(files) > 1)
}

func TestDownloadAnotherBranch(t *testing.T) {
	assert := assert.New(t)
	home := home()
	assert.NoError(project.NewGit(home, "caarlos0/jvm branch:gh-pages").Download())
}

func TestUpdateNonExistentLocalRepo(t *testing.T) {
	assert := assert.New(t)
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	assert.Error(repo.Update())
}

func TestDownloadNonExistentRepo(t *testing.T) {
	var assert = assert.New(t)
	home := home()
	repo := project.NewGit(home, "caarlos0/not-a-real-repo")
	assert.Error(repo.Download())
}

func TestDownloadMalformedRepo(t *testing.T) {
	assert := assert.New(t)
	home := home()
	repo := project.NewGit(home, "doesn-not-exist-really branch:also-nope")
	assert.Error(repo.Download())
}

func TestDownloadMultipleTimes(t *testing.T) {
	assert := assert.New(t)
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	assert.NoError(repo.Download())
	assert.NoError(repo.Download())
	assert.NoError(repo.Update())
}

func TestDownloadFolderNaming(t *testing.T) {
	assert := assert.New(t)
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	assert.Equal(
		home+"/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports",
		repo.Folder(),
	)
}

func TestSubFolder(t *testing.T) {
    assert := assert.New(t)
    home := home()
    repo := project.NewGit(home, "robbyrussell/oh-my-zsh folder:plugins/aws")
    assert.True(strings.HasSuffix(repo.Folder(), "plugins/aws"))
}

func TestMultipleSubFolders(t *testing.T) {
    assert := assert.New(t)
    home := home()
    assert.NoError(project.NewGit(home, "robbyrussell/oh-my-zsh folder:plugins/aws").Download())
    assert.NoError(project.NewGit(home, "robbyrussell/oh-my-zsh folder:plugins/battery").Download())
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
