package project_test

import (
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestLocalProject(t *testing.T) {
	proj := project.NewLocal("/tmp")
	assert.NoError(t, proj.Download())
	assert.NoError(t, proj.Update())
	assert.Equal(t, "/tmp", proj.Folder())
}
