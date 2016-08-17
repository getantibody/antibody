package command

import (
	"fmt"

	"github.com/caarlos0/gohome"
	"github.com/urfave/cli"
)

// Home shows current antibody home folder
var Home = cli.Command{
	Name:    "home",
	Aliases: []string{"prefix", "p"},
	Usage:   "shows the current antibody home folder",
	Action: func(ctx *cli.Context) error {
		fmt.Println(gohome.Cache("antibody") + "/")
		return nil
	},
}
