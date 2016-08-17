package main

import (
	"os"

	"github.com/getantibody/antibody/cmd/antibody/command"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	// defer profile.Start(
	// 	profile.MemProfile,
	// 	profile.CPUProfile,
	// 	profile.NoShutdownHook,
	// 	profile.ProfilePath("."),
	// ).Stop()
	app := cli.NewApp()
	app.Name = "antibody"
	app.Usage = "A faster and simpler antigen written in Golang"
	app.Commands = []cli.Command{
		command.Bundle,
		command.Update,
		command.List,
		command.Home,
		command.Init,
	}
	app.Version = version
	app.Author = "Carlos Alexandro Becker (caarlos0@gmail.com)"
	app.Run(os.Args)
}
