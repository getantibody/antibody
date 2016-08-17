package command

import (
	"fmt"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody/project"
	"github.com/urfave/cli"
)

// List all downloaded bundles
var List = cli.Command{
	Name:  "list",
	Usage: "list all currently downloaded bundles",
	Action: func(ctx *cli.Context) error {
		projects, err := project.List(gohome.Cache("antibody"))
		if err != nil {
			return err
		}
		for _, b := range projects {
			fmt.Println(b)
		}
		return nil
	},
}
