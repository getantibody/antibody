package command

import (
	"fmt"
	"path/filepath"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/project"
	"github.com/urfave/cli"
)

// List all downloaded bundles
var List = cli.Command{
	Name:  "list",
	Usage: "list all currently downloaded bundles",
	Action: func(ctx *cli.Context) error {
		home := antibody.Home()
		projects, err := project.List(home)
		if err != nil {
			return err
		}
		for _, b := range projects {
			fmt.Println(filepath.Join(home, b))
		}
		return nil
	},
}
