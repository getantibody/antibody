package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/akatrevorjay/antibody/antibody/command"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "antibody"
	app.Usage = "A faster and simpler antigen written in Golang"
	app.Commands = []cli.Command{
		command.Bundle, command.Update, command.List, command.Shell, command.Home,
	}
	app.Version = version
	app.Author = "Carlos Alexandro Becker (caarlos0@gmail.com)"
	app.Run(os.Args)
}
