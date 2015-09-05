package internal

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFileCount asserts that the given folder has the amount of
// files/folders inside it.
func AssertFileCount(t *testing.T, total int, folder string) {
	files, _ := ioutil.ReadDir(folder)
	assert.Len(t, files, total)
}
