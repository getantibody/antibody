package bundle_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/stretchr/testify/require"
)

func TestSuccessfullGitBundles(t *testing.T) {
	table := []struct {
		line, result string
	}{
		{
			"caarlos0/jvm",
			"jvm.plugin.zsh\nfpath+=( ",
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
			"caarlos0/jvm kind:fpath",
			"fpath+=( ",
		},
		{
			"docker/cli path:contrib/completion/zsh/_docker",
			"contrib/completion/zsh/_docker",
		},
	}
	for _, row := range table {
		row := row
		t.Run(row.line, func(t *testing.T) {
			t.Parallel()
			home := home()
			result, err := bundle.New(home, row.line).Get()
			require.Contains(t, result, row.result)
			require.NoError(t, err)
		})
	}
}

func TestZshInvalidGitBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "does not exist").Get()
	require.Error(t, err)
}

func TestZshLocalBundle(t *testing.T) {
	home := home()
	require.NoError(t, ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	result, err := bundle.New(home, home).Get()
	require.Contains(t, result, "a.sh")
	require.NoError(t, err)
}

func TestZshInvalidLocalBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda").Get()
	require.Error(t, err)
}

func TestZshBundleWithNoShFiles(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "getantibody/antibody").Get()
	require.NoError(t, err)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	home := home()
	_, err := bundle.New(home, "/asduhasd/asdasda kind:path").Get()
	require.Error(t, err)
}

func TestPathLocalBundle(t *testing.T) {
	home := home()
	require.NoError(t, ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	result, err := bundle.New(home, home+" kind:path").Get()
	require.Equal(t, "export PATH=\""+home+":$PATH\"", result)
	require.NoError(t, err)
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
