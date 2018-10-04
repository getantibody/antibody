package project_test

import (
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/require"
)

func TestLocalProject(t *testing.T) {
	proj := project.NewLocal("/tmp")
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	require.Equal(t, "/tmp", proj.Path())
}
