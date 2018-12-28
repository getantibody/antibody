package project

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	home := home()
	require.NoError(t, New(home, "caarlos0/jvm branch:gh-pages").Download())
	list, err := List(home)
	require.NoError(t, err)
	require.Len(t, list, 1)
}

func TestListEmptyFolder(t *testing.T) {
	home := home()
	list, err := List(home)
	require.NoError(t, err)
	require.Len(t, list, 0)
}

func TestListNonExistentFolder(t *testing.T) {
	list, err := List("/tmp/asdasdadadwhateverwtff")
	require.Error(t, err)
	require.Len(t, list, 0)
}

func TestUpdate(t *testing.T) {
	home := home()
	repo := New(home, "caarlos0/ports")
	require.NoError(t, repo.Download())
	require.NoError(t, repo.Update())
}

func TestUpdateHome(t *testing.T) {
	home := home()
	require.NoError(t, New(home, "caarlos0/jvm").Download())
	require.NoError(t, New(home, "caarlos0/ports").Download())
	require.NoError(t, New(home, "/tmp").Download())
	require.NoError(t, Update(home, runtime.NumCPU()))
}

func TestUpdateNonExistentHome(t *testing.T) {
	require.Error(t, Update("/tmp/asdasdasdasksksksksnopeeeee", runtime.NumCPU()))
}

func TestUpdateHomeWithNoGitProjects(t *testing.T) {
	home := home()
	repo := New(home, "caarlos0/jvm")
	require.NoError(t, repo.Download())
	require.NoError(t, os.RemoveAll(filepath.Join(repo.Path(), ".git")))
	require.Error(t, Update(home, runtime.NumCPU()))
}
