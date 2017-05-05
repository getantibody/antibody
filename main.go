package main

import (
	"github.com/getantibody/antibody/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(version)
}
