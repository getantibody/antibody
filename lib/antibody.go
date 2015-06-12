package antibody

import (
	"fmt"
	"sync"
)

type antibody struct {
	bundles []Bundle
}
type bundleAction func(bundle Bundle)

func NewAntibody(bundles []Bundle) antibody {
	return antibody{bundles}
}

func (a antibody) forEach(fn bundleAction) {
	var wg sync.WaitGroup
	for _, bundle := range a.bundles {
		wg.Add(1)
		go func(bundle Bundle, fn bundleAction, wg *sync.WaitGroup) {
			fn(bundle)
			fmt.Println(bundle.Folder())
			wg.Done()
		}(bundle, fn, &wg)
	}
	wg.Wait()
}

func (a antibody) Download() {
	a.forEach(func(b Bundle) {
		b.Download()
	})
}

func (a antibody) Update() {
	a.forEach(func(b Bundle) {
		b.Update()
	})
}
