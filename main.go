package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func clone(bundle string, home string) string {
	folder := home + strings.Replace(bundle, "/", "-", -1)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--depth", "1", GH+bundle, folder)
		clone.Start()
	}
	return folder
}

func cloneStdin(home string) {
	bundles, _ := ioutil.ReadAll(os.Stdin)
	for _, bundle := range strings.Split(string(bundles), "\n") {
		if bundle != "" {
			fmt.Println(clone(bundle, home))
		}
	}
}

func pull(bundle string, home string) string {
	folder := home + bundle
	pull := exec.Command("git", "-C", folder, "pull", "origin", "master")
	pull.Start()
	return folder
}

func update(home string) {
	bundles, _ := ioutil.ReadDir(home)
	for _, bundle := range bundles {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			fmt.Println(pull(home+bundle.Name(), home))
		}
	}
}

func readStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func main() {
	home := os.Getenv("HOME") + "/.antibody/"
	if readStdin() {
		cloneStdin(home)
	} else {
		if (os.Args[1:][0]) == "update" {
			update(home)
		} else if (os.Args[1:][0]) == "bundle" {
			fmt.Println(clone(os.Args[1:][1], home))
		}
	}
}
