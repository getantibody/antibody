package command

import (
	"github.com/codegangsta/cli"
	"github.com/akatrevorjay/antibody"
	"github.com/akatrevorjay/antibody/bundle"
)

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) {
		antibody.New(bundle.List(antibody.Home())).Update()
	},
}
