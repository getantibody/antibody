package antibody

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/gohome"
)

// Antibody wraps a list of bundles to be processed.
type Antibody struct {
	bundles []Bundle
}

type bundleAction func(bundle Bundle)

// New creates an instance of antibody with the given bundles.
func New(bundles []Bundle) Antibody {
	return Antibody{
		bundles: bundles,
	}
}

// Download the needed bundles.
func (a Antibody) Download() {
	a.forEach(func(b Bundle) {
		b.Download()
	})
}

// Update all bundles.
func (a Antibody) Update() {
	a.forEach(func(b Bundle) {
		b.Update()
	})
}

func (a Antibody) forEach(action bundleAction) {
	var wg sync.WaitGroup
	for _, bundle := range a.bundles {
		wg.Add(1)
		go func(bundle Bundle, action bundleAction) {
			action(bundle)
			for _, sourceable := range bundle.Sourceables() {
				fmt.Println(sourceable)
			}
			wg.Done()
		}(bundle, action)
	}
	wg.Wait()
}

// Home finds the right home folder to use
func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		return gohome.Cache("antibody")
	}
	return home
}
