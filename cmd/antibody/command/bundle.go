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
	static := ctx.Bool("static")
	var bundles []bundle.Bundle
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(ctx.Args()) == 0 {
		entries, err := ioutil.ReadAll(os.Stdin)
		if err != nil || len(entries) == 0 {
			return err
		}
		bundles = bundle.Parse(string(entries), antibody.Home())
	} else {
		bundles = []bundle.Bundle{
			bundle.New(ctx.Args().First(), antibody.Home()),
		}
	}
	if static {
		antibody.NewStatic(bundles).Download()
	} else {
		antibody.New(bundles).Download()
	}
	return nil
}
