package bundle_test

import (
	"io/ioutil"
	"os"
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
		{
			"docker/cli path:contrib/completion/zsh/_docker",
			"contrib/completion/zsh/_docker",
		},
	}
	for _, row := range table {
		t.Run(row.line, func(t *testing.T) {
			home := home()
			result, err := bundle.New(home, row.line).Get()
			assert.Contains(t, result, row.result)
			assert.NoError(t, err)
		})
	}
}

func TestZshInvalidGitBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "doesnt exist").Get()
	assert.Error(t, err)
}

func TestZshLocalBundle(t *testing.T) {
	home := home()
	assert.NoError(t, ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	result, err := bundle.New(home, home).Get()
	assert.Contains(t, result, "a.sh")
	assert.NoError(t, err)
}

func TestZshInvalidLocalBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda").Get()
	assert.Error(t, err)
}

func TestZshBundleWithNoShFiles(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "getantibody/antibody").Get()
	assert.NoError(t, err)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda kind:path").Get()
	assert.Error(t, err)
}

func TestPathLocalBundle(t *testing.T) {
	home := home()
	assert.NoError(t, ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	result, err := bundle.New(home, home+" kind:path").Get()
	assert.Equal(t, "export PATH=\""+home+":$PATH\"", result)
	assert.NoError(t, err)
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
