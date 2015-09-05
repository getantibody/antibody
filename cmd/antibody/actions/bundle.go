package actions

import (
	"io/ioutil"
	"os"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// Bundle download all given bundles (stdin or args)
func Bundle(ctx *cli.Context) {
	if readFromStdin() {
		entries, _ := ioutil.ReadAll(os.Stdin)
		antibody.New(
			bundle.Parse(string(entries), antibody.Home()),
		).Download()
	} else {
		antibody.New([]bundle.Bundle{
			bundle.New(ctx.Args().First(), antibody.Home()),
		}).Download()
	}
}

func readFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
