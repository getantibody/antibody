package antibody

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/event"
)

// Antibody the main thing
type Antibody struct {
	Events chan event.Event
	r      io.Reader
	Home   string
}

// New creates a new Antibody instance with the given parameters
func New(home string, r io.Reader) *Antibody {
	return &Antibody{
		r:      r,
		Events: make(chan event.Event),
		Home:   home,
	}
}

// Bundle processes all given lines and returns the shell content to execute
func (a *Antibody) Bundle() (result string, err error) {
	defer close(a.Events)
	var shs []string
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(a.r)
	for scanner.Scan() {
		wg.Add(1)
		go func(l string) {
			l = strings.TrimSpace(l)
			if l != "" && l[0] != '#' {
				bundle.New(a.Home, l).Get(a.Events)
			} else {
				wg.Done()
			}
		}(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}

	go func() {
		for {
			event := <-a.Events
			if event.Error != nil {
				err = event.Error
			} else {
				shs = append(shs, event.Shell)
			}
			wg.Done()
		}
	}()
	wg.Wait()
	return strings.Join(shs, "\n"), err
}

// Home finds the right home folder to use
func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = gohome.Cache("antibody")
	}
	return home
}
