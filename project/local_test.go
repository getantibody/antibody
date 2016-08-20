package project_test

import (
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestLocalProject(t *testing.T) {
	assert := assert.New(t)
	proj := project.NewLocal("/tmp")
	assert.NoError(proj.Download())
	assert.NoError(proj.Update())
	assert.Equal("/tmp", proj.Folder())
}
