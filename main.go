package main

import (
	"github.com/caarlos0/antibody/lib"
	"os"
)

func main() {
	if antibody.ReadStdin() {
		antibody.ProcessStdin(os.Stdin, antibody.Home())
	} else {
		antibody.ProcessArgs(os.Args[1:], antibody.Home())
	}
}
