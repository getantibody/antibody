package antibody

import (
	"strings"

	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/event"
)

// Antibody the main thing
type Antibody struct {
	Events chan event.Event
	Lines  []string
	Home   string
}

// New creates a new Antibody instance with the given parameters
func New(home string, lines []string) *Antibody {
	return &Antibody{
		Lines:  lines,
		Events: make(chan event.Event),
		Home:   home,
	}
}

// Bundle processes all given lines and returns the shell content to execute
func (a *Antibody) Bundle() (string, error) {
	var shs []string
	var total = len(a.Lines)
	var count int
	done := make(chan bool)

	for _, line := range a.Lines {
		go func(l string) {
			if l != "" && l[0] != '#' {
				bundle.New(a.Home, l).Get(a.Events)
			}
			done <- true
		}(line)
	}
	for {
		select {
		case event := <-a.Events:
			if event.Error != nil {
				return "", event.Error
			}
			shs = append(shs, event.Shell)
		case <-done:
			count++
			if count == total {
				return strings.Join(shs, "\n"), nil
			}
		}
	}
}
