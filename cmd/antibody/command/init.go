package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/getantibody/antibody/shell"
)

// Init prints out the antibody's shell init script
var Init = cli.Command{
	Name:  "init",
	Usage: "Initializes the shell so Antibody can work as expected.",
	Action: func(ctx *cli.Context) {
		fmt.Println(shell.Init())
	},
}
