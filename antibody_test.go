package antibody_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody"
	"github.com/stretchr/testify/assert"
)

func TestAntibody(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
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
	sh, err := antibody.New(home, bundles).Bundle()
	assert.NoError(err)
	files, err := ioutil.ReadDir(home)
	assert.NoError(err)
	assert.Len(files, 3)
	assert.Contains(sh, `export PATH="/tmp:$PATH"`)
	assert.Contains(sh, `export PATH="`+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports:$PATH"`)
	assert.Contains(sh, `export PATH="`+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm:$PATH"`)
	assert.Contains(sh, `source `+home+`/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-zsh-open-pr/open-pr.plugin.zsh`)
}

func TestAntibodyError(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	bundles := []string{"invalid-repo"}
	sh, err := antibody.New(home, bundles).Bundle()
	assert.Error(err)
	assert.Empty(sh)
}

func TestHome(t *testing.T) {
	assert.Contains(t, antibody.Home(), "antibody")
}

func TestHomeFromEnvironmentVariable(t *testing.T) {
	os.Setenv("ANTIBODY_HOME", "/tmp")
	assert.Equal(t, "/tmp", antibody.Home())
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
