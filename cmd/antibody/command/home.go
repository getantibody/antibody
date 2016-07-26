package command

import (
	"fmt"

	"github.com/getantibody/antibody"
	"github.com/urfave/cli"
)

// Home shows current antibody home folder
var Home = cli.Command{
	Name:    "home",
	Aliases: []string{"prefix", "p"},
	Usage:   "shows the current antibody home folder",
	Action: func(ctx *cli.Context) error {
		fmt.Println(antibody.Home())
		return nil
	},
}
