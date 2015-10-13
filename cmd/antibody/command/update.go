package command

import (
	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) {
		antibody.New(bundle.List(antibody.Home())).Update()
	},
}
