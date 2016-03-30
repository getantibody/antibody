package command

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/getantibody/antibody/shell"
)

// Init prints out the antibody's shell init script
var Init = cli.Command{
	Name:  "init",
	Usage: "Initializes the shell so Antibody can work as expected.",
	Action: func(ctx *cli.Context) {
		if shell, err := shell.Init(); err == nil {
			fmt.Println(shell)
		} else {
			log.Fatal(err)
		}
	},
}
