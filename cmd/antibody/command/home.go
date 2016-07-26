package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/getantibody/antibody"
)

// Home shows current antibody home folder
var Home = cli.Command{
	Name:    "home",
	Aliases: []string{"prefix", "p"},
	Usage:   "shows the current antibody home folder",
	Action: func(ctx *cli.Context) {
		fmt.Println(antibody.Home())
	},
}
