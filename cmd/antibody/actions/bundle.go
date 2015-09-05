package actions

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/caarlos0/antibody"
	"github.com/codegangsta/cli"
)

// Bundle download all given bundles (stdin or args)
func Bundle(c *cli.Context) {
	if readFromStdin() {
		processStdin(os.Stdin)
	} else {
		antibody.New(
			[]antibody.Bundle{
				antibody.NewBundle(c.Args().First(), antibody.Home()),
			},
		).Download()
	}
}

func readFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func processStdin(stdin io.Reader) {
	home := antibody.Home()
	entries, _ := ioutil.ReadAll(stdin)
	var bundles []antibody.Bundle
	for _, bundle := range strings.Split(string(entries), "\n") {
		if bundle == "" {
			continue
		}
		bundles = append(bundles, antibody.NewBundle(bundle, home))
	}
	antibody.New(bundles).Download()
}
