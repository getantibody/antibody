package antibodylib

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/bundle"
	"golang.org/x/sync/errgroup"
)

// Antibody the main thing
type Antibody struct {
	r           io.Reader
	parallelism int
	Home        string
}

// New creates a new Antibody instance with the given parameters
func New(home string, r io.Reader, p int) *Antibody {
	return &Antibody{
		r:           r,
		parallelism: p,
		Home:        home,
	}
}

// Bundle processes all given lines and returns the shell content to execute
func (a *Antibody) Bundle() (string, error) {
	var g errgroup.Group
	var shs safeIndexedLines
	var idx int
	var sem = make(chan bool, a.parallelism)
	var scanner = bufio.NewScanner(a.r)
	for scanner.Scan() {
		var line = scanner.Text()
		var index = idx
		idx++
		sem <- true
		g.Go(func() error {
			defer func() {
				<-sem
			}()
			line = strings.TrimSpace(line)
			if line == "" || line[0] == '#' {
				return nil
			}
			lineBundle, berr := bundle.New(a.Home, line)
			if berr != nil {
				return berr
			}
			sh, berr := lineBundle.Get()
			shs.Append(indexedLine{idx: index, line: sh})
			return berr
		})
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	var err = g.Wait()
	return shs.Items().String(), err
}

// Home finds the right home folder to use
func Home() (string, error) {
	if dir := os.Getenv("ANTIBODY_HOME"); dir != "" {
		return dir, nil
	}
	dir, err := os.UserCacheDir()
	return filepath.Join(dir, "antibody"), err
}
