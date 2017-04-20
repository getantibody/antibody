package command

import (
	"fmt"

	"github.com/getantibody/antibody/cmd/shell"
	"github.com/urfave/cli"
)

// Init prints out the antibody's shell init script
var Init = cli.Command{
	Name:  "init",
	Usage: "Initializes the shell so Antibody can work as expected",
	Action: func(ctx *cli.Context) error {
		sh, err := shell.Init()
		fmt.Println(sh)
		return err
	},
}
