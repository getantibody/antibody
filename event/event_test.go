package event_test

import (
	"errors"
	"testing"

	"github.com/getantibody/antibody/event"
	"github.com/stretchr/testify/assert"
)

func TestShellEvent(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("ls", event.Shell("ls").Shell)
}

func TestErrorEvent(t *testing.T) {
	assert := assert.New(t)
	assert.Error(event.Error(errors.New("fake err")).Error)
}
