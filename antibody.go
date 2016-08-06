package antibody

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/bundle"
)

// Antibody wraps a list of bundles to be processed.
type Antibody struct {
	bundles []bundle.Bundle
	static  bool
}

// New creates an instance of antibody with the given bundles.
func New(bundles []bundle.Bundle) Antibody {
	return Antibody{bundles: bundles, static: false}
}

// NewStatic creates an instance of antibody with the given bundles in
// static-loading mode.
func NewStatic(bundles []bundle.Bundle) Antibody {
	return Antibody{bundles: bundles, static: true}
}

// Download the needed bundles.
func (a Antibody) Download() {
	var wg sync.WaitGroup
	for _, b := range a.bundles {
		wg.Add(1)
		go func(b bundle.Bundle) {
			b.Download()
			for _, sourceable := range bundle.Sourceables(b) {
				if a.static {
					fmt.Println("source", sourceable)
				} else {
					fmt.Println(sourceable)
				}
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
