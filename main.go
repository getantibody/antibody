package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func createHome() string {
	fmt.Println("Assuring home exists...")
	home := os.Getenv("HOME") + "/.antibody/"
	os.Mkdir(home, 755)
	return home
}

func clone(home string) string {
	repo := os.Args[1:][0]
	folder := home + strings.Replace(repo, "/", "-", -1)
	fmt.Println("Cloning", repo, "to", folder)
	clone := exec.Command("git", "clone", "--depth", "1", GH+repo, folder)
	clone.Start()
	return folder
}

func source(path string) {
	plugin := path + "/*.plugin.zsh"
	fmt.Println("Sourcing", plugin)
	exec.Command("source", plugin).Start()
}

func main() {
	home := createHome()
	folder := clone(home)
	source(folder)
}
