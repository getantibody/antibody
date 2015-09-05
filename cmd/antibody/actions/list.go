package actions

import (
	"fmt"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// List all installed bundles
func List(c *cli.Context) {
	for _, b := range bundle.List(antibody.Home()) {
		fmt.Println(b.Name())
	}
}
