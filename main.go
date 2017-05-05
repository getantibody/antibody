package main

import "github.com/spf13/cobra/cobra/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
