package antibody_test

import (
	"os"
	"testing"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/caarlos0/antibody/internal"
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
	internal.AssertFileCount(t, 2, home)
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
