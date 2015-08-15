package antibody

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/caarlos0/antibody/doubles"
	"github.com/stretchr/testify/assert"
)

func expectError(t *testing.T) {
	assert.NotNil(t, recover())
}

func assertBundledPlugins(t *testing.T, total int, home string) {
	plugins, _ := ioutil.ReadDir(home)
	assert.Len(t, plugins, total)
}

func TestProcessesArgsDoBundle(t *testing.T) {
	home := doubles.TempHome()
	ProcessArgs([]string{"bundle", "caarlos0/zsh-pg"}, home)
	assertBundledPlugins(t, 1, home)
}

func TestBundleWithNoBundles(t *testing.T) {
	home := doubles.TempHome()
	ProcessArgs([]string{"bundle", ""}, home)
	ProcessArgs([]string{"bundle"}, home)
	assertBundledPlugins(t, 0, home)
}

func TestUpdateWithNoPlugins(t *testing.T) {
	home := doubles.TempHome()
	ProcessArgs([]string{"update"}, home)
	assertBundledPlugins(t, 0, home)
}

func TestVersion(t *testing.T) {
	home := doubles.TempHome()
	ProcessArgs([]string{"version"}, home)
	assertBundledPlugins(t, 0, home)
}

func TestBundleMkdirs(t *testing.T) {
	home := filepath.Join(doubles.TempHome(), "long/folder/which/dont/exist")
	bundle("caarlos0/zsh-pg", home)
	ProcessArgs([]string{"update"}, home)
	assertBundledPlugins(t, 1, home)
}

func TestUpdateWithPlugins(t *testing.T) {
	home := doubles.TempHome()
	bundle("caarlos0/zsh-pg", home)
	ProcessArgs([]string{"update"}, home)
	assertBundledPlugins(t, 1, home)
}

func TestBundlesSinglePlugin(t *testing.T) {
	home := doubles.TempHome()
	bundle("caarlos0/zsh-pg", home)
	assertBundledPlugins(t, 1, home)
}

func TestLoadsDefaultHome(t *testing.T) {
	os.Unsetenv("ANTIBODY_HOME")
	assert.Regexp(t, "/antibody$", Home())
}

func TestLoadsCustomHome(t *testing.T) {
	home := doubles.TempHome()
	assert.Equal(t, home, Home())
}

func TestFailsToBundleInvalidRepos(t *testing.T) {
	home := doubles.TempHome()
	// TODO return an error here
	// defer expectError(t)
	bundle("csadsadp", home)
	assertBundledPlugins(t, 0, home)
}

func TestFailsToProcessInvalidArgs(t *testing.T) {
	home := doubles.TempHome()
	defer expectError(t)
	ProcessArgs([]string{"nope", "caarlos0/zsh-pg"}, home)
	assertBundledPlugins(t, 0, home)
}

func TestReadsStdinIsFalse(t *testing.T) {
	assert.False(t, ReadStdin())
}

func TestReadsStdinIsTrue(t *testing.T) {
	t.SkipNow()
	os.Stdin.Write([]byte("Some STDIN"))
	assert.True(t, ReadStdin())
}

func TestProcessStdin(t *testing.T) {
	home := doubles.TempHome()
	bundles := bytes.NewBufferString("caarlos0/zsh-pg\ncaarlos0/zsh-add-upstream")
	ProcessStdin(bundles, home)
	assertBundledPlugins(t, 2, home)
}

func TestProcessStdinWithEmptyLines(t *testing.T) {
	home := doubles.TempHome()
	bundles := bytes.NewBufferString("\ncaarlos0/zsh-pg\ncaarlos0/zsh-add-upstream\n")
	ProcessStdin(bundles, home)
	assertBundledPlugins(t, 2, home)
}

func TestUpdatesListOfRepos(t *testing.T) {
	home := doubles.TempHome()
	bundle1 := "caarlos0/zsh-pg"
	bundle2 := "caarlos0/zsh-add-upstream"
	NewGitBundle(bundle1, home).Download()
	NewGitBundle(bundle2, home).Download()
	// TODO check amount of updated repos
	update(home)
}

func TestUpdatesBrokenRepo(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-mkc", home)
	bundle.Download()
	os.RemoveAll(bundle.Folder() + "/.git")
	// TODO check amount of updated repos
	update(home)
}
