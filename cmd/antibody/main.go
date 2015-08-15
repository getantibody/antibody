package main

import (
	"os"

	a "github.com/caarlos0/antibody"
)

var version = "master"

func main() {
	if a.ReadStdin() {
		a.ProcessStdin(os.Stdin, a.Home())
	} else {
		a.ProcessArgs(os.Args[1:], a.Home(), version)
	}
}
