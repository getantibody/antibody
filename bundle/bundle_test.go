package bundle_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/event"
	"github.com/stretchr/testify/assert"
)

func TestZshGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "caarlos0/jvm").Get(events)
	assert.Contains((<-events).Shell, "jvm.plugin.zsh")
}

func TestZshInvalidGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "doesnt exists").Get(events)
	assert.Error((<-events).Error)
}

func TestZshLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"/a.sh", []byte("echo 9"), 0644))
	events := make(chan event.Event)
	go bundle.New(home, home).Get(events)
	assert.Contains((<-events).Shell, "a.sh")
}

func TestZshInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "/asduhasd/asdasda").Get(events)
	assert.Error((<-events).Error)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "/asduhasd/asdasda kind:path").Get(events)
	assert.Error((<-events).Error)
}

func TestPathGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "caarlos0/jvm kind:path").Get(events)
	assert.Contains((<-events).Shell, "export PATH=\"")
}

func TestPathLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	events := make(chan event.Event)
	go bundle.New(home, home+" kind:path").Get(events)
	assert.Equal("export PATH=\""+home+":$PATH\"", (<-events).Shell)
}

func TestPathGitBundleWithBranch(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "caarlos0/jvm kind:path branch:gh-pages").Get(events)
	assert.Contains((<-events).Shell, "export PATH=\"")
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home
}
