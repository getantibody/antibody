package antibody

import (
	"bufio"
	"io"
	"os"
	"strings"

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
	var shs []string
	scanner := bufio.NewScanner(a.r)
	for scanner.Scan() {
		l := scanner.Text()
		g.Go(func() error {
			l = strings.TrimSpace(l)
			if l != "" && l[0] != '#' {
				s, err := bundle.New(a.Home, l).Get()
				shs = append(shs, s)
				return err
			}
			return nil
		})
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}
	if err := g.Wait(); err != nil {
		return result, err
	}
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
