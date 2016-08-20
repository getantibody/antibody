package shell_test

import (
	"testing"

	"github.com/getantibody/antibody/cmd/shell"
	"github.com/stretchr/testify/assert"
)

func TestGeneratesInit(t *testing.T) {
	shell := shell.Init()
	assert.NotNil(t, shell)
	assert.NotEmpty(t, shell)
}
