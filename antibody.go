package antibody

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/bundle"
	"golang.org/x/sync/errgroup"
)

// Antibody the main thing
type Antibody struct {
	r    io.Reader
	Home string
}

// New creates a new Antibody instance with the given parameters
func New(home string, r io.Reader) *Antibody {
	return &Antibody{
		r:    r,
		Home: home,
	}
}

// Bundle processes all given lines and returns the shell content to execute
func (a *Antibody) Bundle() (result string, err error) {
	var g errgroup.Group
	var lock sync.Mutex
	var shs []string
	scanner := bufio.NewScanner(a.r)
	for scanner.Scan() {
		l := scanner.Text()
		g.Go(func() error {
			l = strings.TrimSpace(l)

			if l == "" || l[0] == '#' {
				return nil
			}

			s, err := bundle.New(a.Home, l).Get()

			lock.Lock()
			shs = append(shs, s)
			lock.Unlock()

			return err
		})
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}
	err = g.Wait()
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
