package shell_test

import (
	"fmt"
	"testing"

	"github.com/getantibody/antibody/shell"
	"github.com/stretchr/testify/assert"
)

func TestClonesRepo(t *testing.T) {
	shell, err := shell.Init()
	assert.Nil(t, err)
	assert.NotNil(t, shell)
	fmt.Println(shell)
}
