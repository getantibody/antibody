package command

import (
	"io/ioutil"
	"os"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Bundle downloads (if needed) and then sources a given repo
var Bundle = cli.Command{
	Name:   "bundle",
	Usage:  "downloads (if needed) and then sources a given repo",
	Action: doBundle,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "static",
			Usage: "Generates the output in a static-loading compatible way",
		},
	},
}

func doBundle(ctx *cli.Context) error {
	var input string
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(ctx.Args()) == 0 {
		entries, err := ioutil.ReadAll(os.Stdin)
		if err != nil || len(entries) == 0 {
			return err
		}
		input = string(entries)
	} else {
		input = ctx.Args().First()
	}
	if ctx.Bool("static") {
		antibody.NewStatic(bundle.Parse(input, antibody.Home())).Download()
	} else {
		antibody.New(bundle.Parse(input, antibody.Home())).Download()
	}
	return nil
}
