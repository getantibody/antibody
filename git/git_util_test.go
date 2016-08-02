package git_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertFileCount(t *testing.T, total int, home string) {
	files, _ := ioutil.ReadDir(home)
	assert.Len(t, files, total)
}
