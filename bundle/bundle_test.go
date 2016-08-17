package bundle_test

import (
	"io/ioutil"
	"os"
	"strings"
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
	evt := <-events
	assert.True(strings.HasSuffix(evt.Shell, "jvm.plugin.zsh"))
}

func TestZshInvalidGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "doesnt exists").Get(events)
	evt := <-events
	assert.Error(evt.Error)
}

func TestZshLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	events := make(chan event.Event)
	go bundle.New(home, home).Get(events)
	evt := <-events
	assert.True(strings.HasSuffix(evt.Shell, "whatever.sh"))
}

func TestZshInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "/asduhasd/asdasda").Get(events)
	evt := <-events
	assert.Error(evt.Error)
}

func TestPathInvalidLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "/asduhasd/asdasda kind:path").Get(events)
	evt := <-events
	assert.Error(evt.Error)
}

func TestPathGitBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "caarlos0/jvm kind:path").Get(events)
	evt := <-events
	assert.True(strings.HasPrefix(evt.Shell, "export PATH=\""))
}

func TestPathLocalBundle(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(ioutil.WriteFile(home+"whatever.sh", []byte(""), 0644))
	events := make(chan event.Event)
	go bundle.New(home, home+" kind:path").Get(events)
	evt := <-events
	assert.Equal("export PATH=\""+home+":$PATH\"", evt.Shell)
}

func TestPathGitBundleWithBranch(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	events := make(chan event.Event)
	go bundle.New(home, "caarlos0/jvm kind:path branch:gh-pages").Get(events)
	evt := <-events
	assert.True(strings.HasPrefix(evt.Shell, "export PATH=\""))
}

func home() string {
	home, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	return home + "/"
}
