package shell_test

import (
	"testing"

	"github.com/getantibody/antibody/shell"
	"github.com/stretchr/testify/assert"
)

func TestGeneratesInit(t *testing.T) {
	shell, err := shell.Init()
	assert.Nil(t, err)
	assert.NotNil(t, shell)
	assert.NotEmpty(t, shell)
}
