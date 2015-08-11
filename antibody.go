package antibody

import (
	"fmt"
	"sync"
)

// Antibody wraps a list of bundles to be processed.
type Antibody      struct {
	bundles []Bundle
}
type bundleFn func(bundle Bundle)

// NewAntibody creates an instance of antibody with the given bundles.
func NewAntibody(bundles []Bundle) Antibody {
	return Antibody{
		bundles: bundles,
	}
}

func (a Antibody) forEach(fn bundleFn) {
	var wg sync.WaitGroup
	for _, bundle := range a.bundles {
		wg.Add(1)
		go func(bundle Bundle, fn bundleFn) {
			fn(bundle)
			for _, sourceable := range bundle.Sourceables() {
				fmt.Println(sourceable)
			}
			wg.Done()
		}(bundle, fn)
	}
	wg.Wait()
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
