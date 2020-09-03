package project

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadAllKinds(t *testing.T) {
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
		require.NoError(
			t,
			NewGit(home, url).Download(),
			"Repo "+url+" failed to download",
		)
	}
}

func TestDownloadSubmodules(t *testing.T) {
	var home = home()
	var proj = NewGit(home, "fribmendes/geometry branch:master")
	var module = filepath.Join(proj.Path(), "lib/zsh-async")
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	files, err := ioutil.ReadDir(module)
	require.NoError(t, err)
	require.True(t, len(files) > 1)
}

func TestDownloadAnotherBranch(t *testing.T) {
	home := home()
	require.NoError(t, NewGit(home, "caarlos0/jvm branch:gh-pages").Download())
}

func TestUpdateAnotherBranch(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/jvm branch:gh-pages")
	require.NoError(t, repo.Download())
	alreadyClonedRepo := NewClonedGit(home, "https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm")
	require.NoError(t, alreadyClonedRepo.Update())
}

func TestUpdateExistentLocalRepo(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/ports")
	require.NoError(t, repo.Download())
	alreadyClonedRepo := NewClonedGit(home, "https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports")
	require.NoError(t, alreadyClonedRepo.Update())
}

func TestUpdateNonExistentLocalRepo(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/ports")
	require.Error(t, repo.Update())
}

func TestDownloadNonExistentRepo(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/not-a-real-repo")
	require.Error(t, repo.Download())
}

func TestDownloadMalformedRepo(t *testing.T) {
	home := home()
	repo := NewGit(home, "doesn-not-exist-really branch:also-nope")
	require.Error(t, repo.Download())
}

func TestDownloadMultipleTimes(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/ports")
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Update())
}

func TestDownloadFolderNaming(t *testing.T) {
	home := home()
	repo := NewGit(home, "caarlos0/ports")
	require.Equal(
		t,
		home+"/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports",
		repo.Path(),
	)
}

func TestSubFolder(t *testing.T) {
	home := home()
	repo := NewGit(home, "robbyrussell/oh-my-zsh path:plugins/aws")
	require.True(t, strings.HasSuffix(repo.Path(), "plugins/aws"))
}

func TestPath(t *testing.T) {
	home := home()
	repo := NewGit(home, "docker/cli path:contrib/completion/zsh/_docker")
	require.True(t, strings.HasSuffix(repo.Path(), "contrib/completion/zsh/_docker"))
}

func TestMultipleSubFolders(t *testing.T) {
	home := home()
	require.NoError(t, NewGit(home, strings.Join([]string{
		"robbyrussell/oh-my-zsh path:plugins/aws",
		"robbyrussell/oh-my-zsh path:plugins/battery",
	}, "\n")).Download())
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
