package main

import "os"

func main() {
	home := Home()
	if ReadStdin() {
		ProcessStdin(home)
	} else {
		ProcessArgs(os.Args[1:], home)
	}
}
