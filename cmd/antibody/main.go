package main

import (
	"os"

	"github.com/getantibody/antibody/cmd/antibody/command"
	logging "github.com/op/go-logging"
	"github.com/urfave/cli"
)

var version = "master"

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{longfunc}: %{color:bold}%{message} %{color:reset}%{color}@%{shortfile} %{color}#%{level}%{color:reset}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)
	logging.SetLevel(logging.INFO, "")
}

func main() {
	app := cli.NewApp()

	app.Name = "antibody"
	app.Usage = "A faster and simpler antigen written in Golang"
	app.Author = "Carlos Alexandro Becker (caarlos0@gmail.com)"
	app.Version = version

	app.Commands = []cli.Command{
		command.Bundle,
		command.Update,
		command.List,
		command.Home,
		command.Init,
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Be more verbose",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("verbose") {
			logging.SetLevel(logging.DEBUG, "")
		}
		return nil
	}

	app.Run(os.Args)
}
