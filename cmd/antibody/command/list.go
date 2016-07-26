package command

import (
	"fmt"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/urfave/cli"
)

// List all downloaded bundles
var List = cli.Command{
	Name:  "list",
	Usage: "list all currently downloaded bundles",
	Action: func(ctx *cli.Context) error {
		for _, b := range bundle.List(antibody.Home()) {
			fmt.Println(b.Name())
		}
		return nil
	},
}
