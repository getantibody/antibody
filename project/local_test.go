package project

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalProject(t *testing.T) {
	proj, err := NewLocal("/tmp")
	require.NoError(t, err)
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	require.Equal(t, "/tmp", proj.Path())
}

func TestLocalProjectRelativeToHome(t *testing.T) {
	proj, err := NewLocal("~/tmp")
	require.NoError(t, err)
	home, err := os.UserHomeDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, "tmp"), proj.Path())
}
