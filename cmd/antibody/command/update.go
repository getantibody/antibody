package command

import "github.com/urfave/cli"

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) error {
		// TODO
		// antibody.New(bundle.List(antibody.Home())).Update()
		return nil
	},
}
