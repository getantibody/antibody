package actions

import (
	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// Update all installed bundles
func Update(ctx *cli.Context) {
	antibody.New(bundle.List(antibody.Home())).Update()
}
