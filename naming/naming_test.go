package naming_test

import (
	"testing"

	"github.com/getantibody/antibody/naming"
	"github.com/stretchr/testify/assert"
)

func TestNaming(t *testing.T) {
	data := []struct {
		url, folder string
	}{
		{"http://google.com", "http-COLON--SLASH--SLASH-google.com"},
		{"git@github.com:getantibody/antibody.git", "git-AT-github.com-COLON-getantibody-SLASH-antibody.git"},
	}

	assert := assert.New(t)
	for _, d := range data {
		assert.Equal(d.folder, naming.URLToFolder(d.url))
		assert.Equal(d.url, naming.FolderToURL(d.folder))
	}
}
