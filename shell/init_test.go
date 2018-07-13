package shell_test

import (
	"testing"

	"github.com/getantibody/antibody/shell"
	"github.com/stretchr/testify/require"
)

func TestGeneratesInit(t *testing.T) {
	shell, err := shell.Init()
	require.NoError(t, err)
	require.NotEmpty(t, shell)
}
