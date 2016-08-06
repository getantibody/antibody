package antibody_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestBundleAndUpdate(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	a := antibody.New([]bundle.Bundle{
		bundle.New("caarlos0/zsh-pg", home),
		bundle.New("caarlos0/zsh-open-pr", home),
	})
	a.Download()
	a.Update()
	files, _ := ioutil.ReadDir(home)
	assert.Len(t, files, 2)
}

func TestBundleAndUpdateStatic(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	a := antibody.NewStatic([]bundle.Bundle{
		bundle.New("caarlos0/zsh-pg", home),
		bundle.New("caarlos0/zsh-open-pr", home),
	})
	a.Download()
	a.Update()
	files, _ := ioutil.ReadDir(home)
	assert.Len(t, files, 2)
}

func TestCustomHome(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	assert.Equal(t, home, antibody.Home())
}

func TestDefaultHome(t *testing.T) {
	os.Unsetenv("ANTIBODY_HOME")
	assert.NotEmpty(t, antibody.Home())
}
