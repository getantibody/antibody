package actions

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// Bundle download all given bundles (stdin or args)
func Bundle(c *cli.Context) {
	if readFromStdin() {
		processStdin(os.Stdin)
	} else {
		antibody.New([]bundle.Bundle{
			bundle.New(c.Args().First(), antibody.Home()),
		}).Download()
	}
}

func readFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func processStdin(stdin io.Reader) {
	home := antibody.Home()
	entries, _ := ioutil.ReadAll(stdin)
	var bundles []bundle.Bundle
	for _, b := range strings.Split(string(entries), "\n") {
		if b == "" {
			continue
		}
		bundles = append(bundles, bundle.New(b, home))
	}
	antibody.New(bundles).Download()
}
