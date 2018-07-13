package project_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getantibody/antibody/project"
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
			project.NewGit(home, url).Download(),
			"Repo "+url+" failed to download",
		)
	}
}

func TestDownloadSubmodules(t *testing.T) {
	var home = home()
	var proj = project.NewGit(home, "fribmendes/geometry")
	var module = filepath.Join(proj.Folder(), "lib/zsh-async")
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	files, err := ioutil.ReadDir(module)
	require.NoError(t, err)
	require.True(t, len(files) > 1)
}

func TestDownloadAnotherBranch(t *testing.T) {
	home := home()
	require.NoError(t, project.NewGit(home, "caarlos0/jvm branch:gh-pages").Download())
}

func TestUpdateNonExistentLocalRepo(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	require.Error(t, repo.Update())
}

func TestDownloadNonExistentRepo(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "caarlos0/not-a-real-repo")
	require.Error(t, repo.Download())
}

func TestDownloadMalformedRepo(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "doesn-not-exist-really branch:also-nope")
	require.Error(t, repo.Download())
}

func TestDownloadMultipleTimes(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Update())
}

func TestDownloadFolderNaming(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "caarlos0/ports")
	require.Equal(
		t,
		home+"/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports",
		repo.Folder(),
	)
}

func TestSubFolder(t *testing.T) {
	home := home()
	repo := project.NewGit(home, "robbyrussell/oh-my-zsh folder:plugins/aws")
	require.True(t, strings.HasSuffix(repo.Folder(), "plugins/aws"))
}

func TestMultipleSubFolders(t *testing.T) {
	home := home()
	require.NoError(t, project.NewGit(home, strings.Join([]string{
		"robbyrussell/oh-my-zsh folder:plugins/aws",
		"robbyrussell/oh-my-zsh folder:plugins/battery",
	}, "\n")).Download())
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
