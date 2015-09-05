package main

import (
	"os"

	"github.com/caarlos0/antibody/cmd/antibody/actions"
	"github.com/codegangsta/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "antibody"
	app.Usage = "antibody is a faster and leaner version of antigen"
	app.Commands = []cli.Command{
		{
			Name:   "bundle",
			Usage:  "bundle one or more bundles",
			Action: actions.Bundle,
		}, {
			Name:   "update",
			Usage:  "updates all previously bundled commands",
			Action: actions.Update,
		}, {
			Name:   "list",
			Usage:  "list all currently installed bundles",
			Action: actions.List,
		},
	}
	app.Version = version
	app.Run(os.Args)
}
