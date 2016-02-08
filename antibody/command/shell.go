package command

import (
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/getantibody/antibody"
)

// Init outputs hooks meant to integrate with your shell.
var Shell = cli.Command{
	Name: "shell",
	//Aliases: []string{"shell_init"},
	Usage: "Generates injection wrapper for your shell. You source the output like so: eval $(antibody shell_init -)",
	Action: func(ctx *cli.Context) {
		var template_name = "shell_init.zsh"

		// use asset data from go-binfmt
		data, err := antibody.Asset(template_name)
		if err != nil {
			// Asset was not found.
			log.Fatalf("Could not retrieve hook template %s: %s",
				template_name, err)
		}

		var sdata = string(data[:])
		var hook_template = template.Must(template.New(template_name).Parse(sdata))

		// Return templated shell hook function.
		err = hook_template.Execute(os.Stdout, ctx.App)
		if err != nil {
			log.Fatalf("Could not generate hook from template: %s", err)
		}
	},
}
