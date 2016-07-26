package command

import (
	"github.com/urfave/cli"
	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
)

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) {
		antibody.New(bundle.List(antibody.Home())).Update()
	},
}
