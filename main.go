package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func clone(home string) string {
	repo := os.Args[1:][0]
	folder := home + strings.Replace(repo, "/", "-", -1)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		// fmt.Println("Cloning", repo)
		clone := exec.Command("git", "clone", "--depth", "1", GH+repo, folder)
		clone.Start()
	}
	return folder
}

func main() {
	home := os.Getenv("HOME") + "/.antibody/"
	folder := clone(home)
	fmt.Println(folder)
}
