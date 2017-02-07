package project_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(project.New(home, "caarlos0/jvm branch:gh-pages").Download())
	list, err := project.List(home)
	assert.NoError(err)
	assert.Len(list, 1)
}

func TestListEmptyFolder(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	list, err := project.List(home)
	assert.NoError(err)
	assert.Len(list, 0)
}

func TestListNonExistentFolder(t *testing.T) {
	assert := assert.New(t)
	list, err := project.List("/tmp/asdasdadadwhateverwtff")
	assert.Error(err)
	assert.Len(list, 0)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.New(home, "caarlos0/ports")
	assert.NoError(repo.Download())
	assert.NoError(repo.Update())
}

func TestUpdateHome(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(project.New(home, "caarlos0/jvm").Download())
	assert.NoError(project.New(home, "caarlos0/ports").Download())
	assert.NoError(project.New(home, "/tmp").Download())
	assert.NoError(project.Update(home))
}

func TestUpdateNonExistentHome(t *testing.T) {
	assert.Error(t, project.Update("/tmp/asdasdasdasksksksksnopeeeee"))
}

func TestUpdateHomeWithNoGitProjects(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	repo := project.New(home, "caarlos0/jvm")
	assert.NoError(repo.Download())
	os.RemoveAll(filepath.Join(repo.Folder(), ".git"))
	assert.Error(project.Update(home))
}
