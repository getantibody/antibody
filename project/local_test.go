package project

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalProject(t *testing.T) {
	proj := NewLocal("/tmp")
	require.NoError(t, proj.Download())
	require.NoError(t, proj.Update())
	require.Equal(t, "/tmp", proj.Path())
}

func TestLocalProjectRelativeToHome(t *testing.T) {
	proj := NewLocal("~/tmp")
	usr, _ := user.Current()
	require.Equal(t, usr.HomeDir+"/tmp", proj.Path())
}
