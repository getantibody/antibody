package antibody

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/bundle"
)

type print func(s string)

func sourcedPrint(s string) {
	fmt.Println(s)
}

func staticPrint(s string) {
	fmt.Println("source", s)
}

// Antibody wraps a list of bundles to be processed.
type Antibody struct {
	bundles []bundle.Bundle
	print
}

// New creates an instance of antibody with the given bundles.
func New(bundles []bundle.Bundle) Antibody {
	return Antibody{bundles, sourcedPrint}
}

// NewStatic creates an instance of antibody with the given bundles in
// static-loading mode.
func NewStatic(bundles []bundle.Bundle) Antibody {
	return Antibody{bundles, staticPrint}
}

// Download the needed bundles.
func (a Antibody) Download() {
	var wg sync.WaitGroup
	for _, b := range a.bundles {
		wg.Add(1)
		go func(b bundle.Bundle) {
			b.Download()
			for _, sourceable := range bundle.Sourceables(b) {
				a.print(sourceable)
			}
			wg.Done()
		}(b)
	}
	wg.Wait()
}

// Update all bundles.
func (a Antibody) Update() {
	var wg sync.WaitGroup
	fmt.Println("Updating...")
	for _, b := range a.bundles {
		wg.Add(1)
		go func(b bundle.Bundle) {
			b.Update()
			wg.Done()
		}(b)
	}
	wg.Wait()
}

// Home finds the right home folder to use
func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = gohome.Cache("antibody")
	}
	return home
}
