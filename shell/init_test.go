package shell

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratesInit(t *testing.T) {
	shell, err := Init()
	require.NoError(t, err)
	require.NotEmpty(t, shell)
}
