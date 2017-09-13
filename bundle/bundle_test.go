package bundle_test

import (
	"io/ioutil"
	"path"
	"os"
	"strings"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfullGitBundles(t *testing.T) {
	table := []struct {
		line, result string
	}{
		{
			"caarlos0/jvm",
			"jvm.plugin.zsh",
		},
		{
			"caarlos0/jvm kind:path",
			"export PATH=\"",
		},
		{
			"caarlos0/jvm kind:path branch:gh-pages",
			"export PATH=\"",
		},
		{
			"caarlos0/jvm kind:dummy",
			"",
		},
	}
	for _, row := range table {
		t.Run(row.line, func(t *testing.T) {
			assert := assert.New(t)
			home := home()
			result, err := bundle.New(home, row.line).Get()
			assert.Contains(result, row.result)
			assert.NoError(err)
		})
	}
}

func TestPreGitBundle(t *testing.T) {
	assert := assert.New(t)
	homeDir := home()

	preDir := home()
	os.Setenv("ANTIBODY_BUNDLE_PRE", preDir)
	bundlePre := path.Join(preDir, "https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm")
	os.Mkdir(bundlePre, 0777)
	os.OpenFile(path.Join(bundlePre, "test.zsh"), os.O_RDONLY|os.O_CREATE, 0666)

	result, err := bundle.New(homeDir, "caarlos0/jvm").Get()
	lines := strings.Split(result, "\n")
	assert.Equal(len(lines), 2)
	assert.Contains(lines[0], preDir)
	assert.Contains(lines[0], bundlePre)
	assert.Contains(lines[0], "test.zsh")
	assert.Contains(lines[1], "jvm.plugin.zsh")
	assert.NoError(err)
	os.Unsetenv("ANTIBODY_BUNDLE_PRE")
}

func TestPostGitBundle(t *testing.T) {
	assert := assert.New(t)
	homeDir := home()

	postDir := home()
	os.Setenv("ANTIBODY_BUNDLE_POST", postDir)
	bundlePost := path.Join(postDir, "https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm")
	os.Mkdir(bundlePost, 0777)
	os.OpenFile(path.Join(bundlePost, "test.zsh"), os.O_RDONLY|os.O_CREATE, 0666)

	result, err := bundle.New(homeDir, "caarlos0/jvm").Get()
	lines := strings.Split(result, "\n")
	assert.Equal(len(lines), 2)
	assert.Contains(lines[0], "jvm.plugin.zsh")
	assert.Contains(lines[1], postDir)
	assert.Contains(lines[1], bundlePost)
	assert.Contains(lines[1], "test.zsh")
	assert.NoError(err)
	os.Unsetenv("ANTIBODY_BUNDLE_POST")
}

func TestZshInvalidGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	_, err := bundle.New(home, "doesnt exist").Get()
	assert.Error(err)
}

func TestZshLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	assert.NoError(ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	result, err := bundle.New(home, home).Get()
	assert.Contains(result, "a.sh")
	assert.NoError(err)
}

func TestZshInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda").Get()
	assert.Error(err)
}

func TestZshBundleWithNoShFiles(t *testing.T) {
	assert := assert.New(t)
	home := home()
	_, err := bundle.New(home, "getantibody/antibody").Get()
	assert.NoError(err)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda kind:path").Get()
	assert.Error(err)
}

func TestPathLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	assert.NoError(ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	result, err := bundle.New(home, home+" kind:path").Get()
	assert.Equal("export PATH=\""+home+":$PATH\"", result)
	assert.NoError(err)
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
