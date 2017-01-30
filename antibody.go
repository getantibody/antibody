package antibody

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/bradfitz/slice"
	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/bundle"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("antibody")

// Antibody the main thing
type Antibody struct {
	r    io.Reader
	Home string
}

// Result is a pipeline element messenger object
type Result struct {
	idx  int
	line string
	err  error
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
	file := a.r

	inputLines := make(chan Result)
	results := make(chan Result)

	// I think we need a wait group, not sure.
	wg := new(sync.WaitGroup)

	// workers
	for w := 1; w <= 8; w++ {
		wg.Add(1)
		go func() {
			// Decreasing internal counter for wait-group as soon as goroutine finishes
			defer wg.Done()

			for res := range inputLines {
				log.Debugf("Bundling: %s", res.line)
				res.line = strings.TrimSpace(res.line)

				if res.line == "" || res.line[0] == '#' {
					continue
				}

				val, err := bundle.New(a.Home, res.line).Get()
				res.line = val
				res.err = err

				if err != nil {
					log.Errorf("Error processing line (%s): %s", err, res.line)
				}

				results <- res
			}
		}()
	}

	// source
	go func() {
		log.Debugf("Reading bundles")

		scan := bufio.NewScanner(file)

		idx := 0
		for scan.Scan() {
			line := scan.Text()
			line = strings.TrimSpace(line)

			inputLines <- Result{idx, line, nil}
			idx++
		}

		err := scan.Err()
		if err != nil {
			log.Fatalf("Error reading bundles: %s", err)
		}

		close(inputLines)

		log.Debugf("Done reading bundles")
	}()

	// waiter
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect
	var allResults []Result
	for res := range results {
		allResults = append(allResults, res)
	}

	// sort by original idx
	slice.Sort(allResults[:], func(i, j int) bool {
		return allResults[i].idx < allResults[j].idx
	})

	// get values
	var sources []string
	for _, res := range allResults {
		sources = append(sources, res.line)

		// Does Go have chained exceptions? Course not. Oh well.
		if res.err != nil {
			err = res.err
		}
	}

	// coallesce
	return strings.Join(sources, "\n"), err
}

// Home finds the right home folder to use
func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = gohome.Cache("antibody")
	}
	return home
}
