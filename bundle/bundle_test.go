package bundle

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

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
			home := home(t)
			bundle, err := New(home, row.line)
			require.NoError(t, err)
			result, err := bundle.Get()
			require.Contains(t, result, row.result)
			require.NoError(t, err)
		})
	}
}

func TestZshInvalidGitBundle(t *testing.T) {
	home := home(t)
	bundle, err := New(home, "does not exist")
	require.NoError(t, err)
	_, err = bundle.Get()
	require.Error(t, err)
}

func TestZshLocalBundle(t *testing.T) {
	home := home(t)
	require.NoError(t, ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	bundle, err := New(home, home)
	require.NoError(t, err)
	result, err := bundle.Get()
	require.Contains(t, result, "a.sh")
	require.NoError(t, err)
}

func TestZshInvalidLocalBundle(t *testing.T) {
	home := home(t)
	bundle, err := New(home, "/asduhasd/asdasda")
	require.NoError(t, err)
	_, err = bundle.Get()
	require.Error(t, err)
}

func TestZshBundleWithNoShFiles(t *testing.T) {
	home := home(t)
	bundle, err := New(home, "getantibody/antibody")
	require.NoError(t, err)
	_, err = bundle.Get()
	require.NoError(t, err)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	home := home(t)
	bundle, err := New(home, "/asduhasd/asdasda kind:path")
	require.NoError(t, err)
	_, err = bundle.Get()
	require.Error(t, err)
}

func TestPathLocalBundle(t *testing.T) {
	home := home(t)
	require.NoError(t, ioutil.WriteFile(filepath.Join(home, "whatever.sh"), []byte(""), 0644))
	bundle, err := New(home, home+" kind:path")
	require.NoError(t, err)
	result, err := bundle.Get()
	require.NoError(t, err)
	require.Equal(t, "export PATH=\""+home+":$PATH\"", result)
	require.NoError(t, err)
}

func home(t *testing.T) string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	require.NoError(t, err)
	return home
}
