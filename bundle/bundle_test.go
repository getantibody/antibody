package bundle_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/stretchr/testify/assert"
)

func TestZshGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	result, err := bundle.New(home, "caarlos0/jvm").Get()
	assert.Contains(result, "jvm.plugin.zsh")
	assert.NoError(err)
}

func TestZshInvalidGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	_, err := bundle.New(home, "doesnt exist").Get()
	assert.Error(err)
}

func TestZshLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	result, err := bundle.New(home, home).Get()
	assert.Contains(result, "a.sh")
	assert.NoError(err)
}

func TestZshInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	_, err := bundle.New(home, "/asduhasd/asdasda").Get()
	assert.Error(err)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	_, err := bundle.New(home, "/asduhasd/asdasda kind:path").Get()
	assert.Error(err)
}

func TestPathGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	result, err := bundle.New(home, "caarlos0/jvm kind:path").Get()
	assert.Contains(result, "export PATH=\"")
	assert.NoError(err)
}

func TestPathLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	result, err := bundle.New(home, home+" kind:path").Get()
	assert.Equal("export PATH=\""+home+":$PATH\"", result)
	assert.NoError(err)
}

func TestPathGitBundleWithBranch(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	result, err := bundle.New(home, "caarlos0/jvm kind:path branch:gh-pages").Get()
	assert.Contains(result, "export PATH=\"")
	assert.NoError(err)
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
