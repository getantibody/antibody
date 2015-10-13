package command

import (
	"io/ioutil"
	"os"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

var Bundle = cli.Command{
	Name:   "bundle",
	Usage:  "bundle one or more bundles",
	Action: doBundle,
}

func doBundle(ctx *cli.Context) {
	if readFromStdin() {
		entries, _ := ioutil.ReadAll(os.Stdin)
		antibody.New(
			bundle.Parse(string(entries), antibody.Home()),
		).Download()
	} else {
		antibody.New(
			[]bundle.Bundle{bundle.New(ctx.Args().First(), antibody.Home())},
		).Download()
	}
}

func readFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
