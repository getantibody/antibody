package main

import "os"

func main() {
	if ReadStdin() {
		ProcessStdin(os.Stdin, Home())
	} else {
		ProcessArgs(os.Args[1:], Home())
	}
}
