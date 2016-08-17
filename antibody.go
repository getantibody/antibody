package antibody

import (
	"strings"

	"github.com/getantibody/antibody/bundle"
	"github.com/getantibody/antibody/event"
)

type Antibody struct {
	Events  chan event.Event
	Bundles []string
	Home    string
}

func New(home string, bundles []string) *Antibody {
	return &Antibody{
		Bundles: bundles,
		Events:  make(chan event.Event, len(bundles)),
		Home:    home,
	}
}

func (a *Antibody) Bundle() (sh string, err error) {
	var shs []string
	var total = len(a.Bundles)
	var count int
	done := make(chan bool)

	for _, sh := range a.Bundles {
		go func(s string) {
			bundle.New(a.Home, s).Get(a.Events)
			done <- true
		}(sh)
	}
	for {
		select {
		case event := <-a.Events:
			if event.Error != nil {
				return "", err
			}
			shs = append(shs, event.Shell)
		case <-done:
			count++
			if count == total {
				return strings.Join(shs, "\n"), nil
			}
		}
	}
}
