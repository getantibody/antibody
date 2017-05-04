package shell_test

import (
	"testing"

	"github.com/getantibody/antibody/cmd/shell"
	"github.com/stretchr/testify/assert"
)

func TestGeneratesInit(t *testing.T) {
	shell, err := shell.Init()
	assert.NoError(t, err)
	assert.NotEmpty(t, shell)
}
