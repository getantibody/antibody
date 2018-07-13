package project_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	home := home()
	assert.NoError(t, project.New(home, "caarlos0/jvm branch:gh-pages").Download())
	list, err := project.List(home)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestListEmptyFolder(t *testing.T) {
	home := home()
	list, err := project.List(home)
	assert.NoError(t, err)
	assert.Len(t, list, 0)
}

func TestListNonExistentFolder(t *testing.T) {
	list, err := project.List("/tmp/asdasdadadwhateverwtff")
	assert.Error(t, err)
	assert.Len(t, list, 0)
}

func TestUpdate(t *testing.T) {
	home := home()
	repo := project.New(home, "caarlos0/ports")
	assert.NoError(t, repo.Download())
	assert.NoError(t, repo.Update())
}

func TestUpdateHome(t *testing.T) {
	home := home()
	assert.NoError(t, project.New(home, "caarlos0/jvm").Download())
	assert.NoError(t, project.New(home, "caarlos0/ports").Download())
	assert.NoError(t, project.New(home, "/tmp").Download())
	assert.NoError(t, project.Update(home, runtime.NumCPU()))
}

func TestUpdateNonExistentHome(t *testing.T) {
	assert.Error(t, project.Update("/tmp/asdasdasdasksksksksnopeeeee", runtime.NumCPU()))
}

func TestUpdateHomeWithNoGitProjects(t *testing.T) {
	home := home()
	repo := project.New(home, "caarlos0/jvm")
	assert.NoError(t, repo.Download())
	assert.NoError(t, os.RemoveAll(filepath.Join(repo.Folders()[0], ".git")))
	assert.Error(t, project.Update(home, runtime.NumCPU()))
}
