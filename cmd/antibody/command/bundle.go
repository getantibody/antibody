package command

import (
	"io/ioutil"
	"os"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/codegangsta/cli"
)

// Bundle downloads (if needed) and then sources a given repo
var Bundle = cli.Command{
	Name:   "bundle",
	Usage:  "downloads (if needed) and then sources a given repo",
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
