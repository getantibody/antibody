package antibody

// Go will execute this to create the necessary bindata as a part of 'go generate ./...'
//XX go:generate go-bindata -o ./internal/antibody/bindata.go -pkg $GOPACKAGE -prefix data/ data/
//go:generate go-bindata -o ./bindata.go -pkg $GOPACKAGE -prefix data/ data/

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/gohome"
	"github.com/akatrevorjay/antibody/bundle"
)

// Antibody wraps a list of bundles to be processed.
type Antibody struct {
	bundles []bundle.Bundle
}

type bundleAction func(b bundle.Bundle)

// New creates an instance of antibody with the given bundles.
func New(bundles []bundle.Bundle) Antibody {
	return Antibody{bundles: bundles}
}

// Download the needed bundles.
func (a Antibody) Download() {
	a.forEach(func(b bundle.Bundle) {
		b.Download()
	})
}

// Update all bundles.
func (a Antibody) Update() {
	a.forEach(func(b bundle.Bundle) {
		b.Update()
	})
}

func (a Antibody) forEach(action bundleAction) {
	var wg sync.WaitGroup
	for _, b := range a.bundles {
		wg.Add(1)
		go func(b bundle.Bundle, action bundleAction) {
			action(b)
			for _, sourceable := range b.Sourceables() {
				fmt.Println(sourceable)
			}
			wg.Done()
		}(b, action)
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
