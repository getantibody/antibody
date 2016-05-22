package bundle_test

import (
	"os"
	"testing"

	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestDownloadFolder(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New(home+"/nope", home)
	b.Download()
	assert.Empty(t, bundle.Sourceables(b))
}

func TestUpdateFolder(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New(home+"/nope", home)
	b.Update()
	assert.Empty(t, bundle.Sourceables(b))
}

func TestGetFolderName(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	b := bundle.New(home+"nope", home)
	assert.Equal(t, "nope", b.Name())
}
