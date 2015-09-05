package antibody_test

import (
	"os"
	"testing"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestBundleSourceables(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := antibody.NewBundle("caarlos0/zsh-pg", home)
	b.Download()
	assert.NotEmpty(t, b.Sourceables())
}

func TestSourceablesWithoutDownload(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := antibody.NewBundle("caarlos0/zsh-pg", home)
	assert.Empty(t, b.Sourceables())
}
