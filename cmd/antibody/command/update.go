package command

import (
	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/project"
	"github.com/urfave/cli"
)

// Update all previously bundled bundles
var Update = cli.Command{
	Name:  "update",
	Usage: "updates all previously bundled commands",
	Action: func(ctx *cli.Context) error {
		return project.Update(gohome.Cache("antibody"))
	},
}
