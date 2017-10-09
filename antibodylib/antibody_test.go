package antibodylib_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/stretchr/testify/assert"
)

func TestAntibody(t *testing.T) {
	assert := assert.New(t)
	home := home()
	bundles := []string{
		"# comments also are allowed",
		"caarlos0/ports kind:path # comment at the end of the line",
		"caarlos0/jvm kind:path branch:gh-pages",
		"caarlos0/zsh-open-pr     kind:zsh",
		"",
		"        ",
		"  # trick play",
		"/tmp kind:path",
	}
	sh, err := antibodylib.New(
		home,
		bytes.NewBufferString(strings.Join(bundles, "\n")),
		runtime.NumCPU(),
	).Bundle()
	assert.NoError(err)
	files, err := ioutil.ReadDir(home)
	assert.NoError(err)
	assert.Len(files, 3)
	assert.Contains(sh, `export PATH="/tmp:$PATH"`)
	assert.Contains(sh, `export PATH="`+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports:$PATH"`)
	assert.Contains(sh, `export PATH="`+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm:$PATH"`)
	assert.Contains(sh, `source `+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-zsh-open-pr/git-open-pr.plugin.zsh`)
}

func TestAntibodyError(t *testing.T) {
	assert := assert.New(t)
	home := home()
	bundles := bytes.NewBufferString("invalid-repo")
	sh, err := antibodylib.New(home, bundles, runtime.NumCPU()).Bundle()
	assert.Error(err)
	assert.Empty(sh)
}

func TestMultipleRepositories(t *testing.T) {
	var assert = assert.New(t)
	home := home()
	bundles := []string{
		"# this block is in alphabetic order",
		"caarlos0/git-add-remote kind:path",
		"caarlos0/jvm",
		"caarlos0/ports kind:path",
		"caarlos0/zsh-git-fetch-merge kind:path",
		"caarlos0/zsh-git-sync kind:path",
		"caarlos0/zsh-mkc",
		"caarlos0/zsh-open-pr kind:path",
		"mafredri/zsh-async",
		"rupa/z",
		"Tarrasch/zsh-bd",
		"wbinglee/zsh-wakatime",
		"zsh-users/zsh-completions",
		"zsh-users/zsh-autosuggestions",
		"",
		"# these should be at last!",
		"sindresorhus/pure",
		"zsh-users/zsh-syntax-highlighting",
		"zsh-users/zsh-history-substring-search",
	}
	sh, err := antibodylib.New(
		home,
		bytes.NewBufferString(strings.Join(bundles, "\n")),
		runtime.NumCPU(),
	).Bundle()
	assert.NoError(err)
	assert.Len(strings.Split(sh, "\n"), 16)
}

func TestHome(t *testing.T) {
	assert.Contains(t, antibodylib.Home(), "antibody")
}

func TestHomeFromEnvironmentVariable(t *testing.T) {
	assert.NoError(t, os.Setenv("ANTIBODY_HOME", "/tmp"))
	assert.Equal(t, "/tmp", antibodylib.Home())
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
