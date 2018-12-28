package project

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalProject(t *testing.T) {
	proj := NewLocal("/tmp")
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	require.Equal(t, "/tmp", proj.Path())
}
