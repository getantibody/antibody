package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/akatrevorjay/antibody"
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
