package main

import (
	"os"

	"github.com/caarlos0/antibody/cmd/antibody/command"
	"github.com/codegangsta/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "antibody"
	app.Usage = "A faster and simpler antigen written in Golang"
	app.Commands = []cli.Command{
		command.Bundle, command.Update, command.List, command.Home,
	}
	app.Version = version
	app.Author = "Carlos Alexandro Becker (caarlos0@gmail.com)"
	app.Run(os.Args)
}
