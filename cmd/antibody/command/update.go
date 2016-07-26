package command

import (
	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/urfave/cli"
)

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) error {
		antibody.New(bundle.List(antibody.Home())).Update()
		return nil
	},
}
