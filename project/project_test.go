package project_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	home := home()
	require.NoError(t, project.New(home, "caarlos0/jvm branch:gh-pages").Download())
	list, err := project.List(home)
	require.NoError(t, err)
	require.Len(t, list, 1)
}

func TestListEmptyFolder(t *testing.T) {
	home := home()
	list, err := project.List(home)
	require.NoError(t, err)
	require.Len(t, list, 0)
}

func TestListNonExistentFolder(t *testing.T) {
	list, err := project.List("/tmp/asdasdadadwhateverwtff")
	require.Error(t, err)
	require.Len(t, list, 0)
}

func TestUpdate(t *testing.T) {
	home := home()
	repo := project.New(home, "caarlos0/ports")
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Update())
}

func TestUpdateHome(t *testing.T) {
	home := home()
	require.NoError(t, project.New(home, "caarlos0/jvm").Download())
	require.NoError(t, project.New(home, "caarlos0/ports").Download())
	require.NoError(t, project.New(home, "/tmp").Download())
	require.NoError(t, project.Update(home, runtime.NumCPU()))
}

func TestUpdateNonExistentHome(t *testing.T) {
	require.Error(t, project.Update("/tmp/asdasdasdasksksksksnopeeeee", runtime.NumCPU()))
}

func TestUpdateHomeWithNoGitProjects(t *testing.T) {
	home := home()
	repo := project.New(home, "caarlos0/jvm")
	require.NoError(t, repo.Download())
	require.NoError(t, os.RemoveAll(filepath.Join(repo.Path(), ".git")))
	require.Error(t, project.Update(home, runtime.NumCPU()))
}
