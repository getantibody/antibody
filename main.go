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

func main() {
	home := os.Getenv("HOME") + "/.antibody/"
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bundles, _ := ioutil.ReadAll(os.Stdin)
		for _, bundle := range strings.Split(string(bundles), "\n") {
			if bundle != "" {
				fmt.Println(clone(bundle, home))
			}
		}
	} else {
		fmt.Println(clone(os.Args[1:][0], home))
	}
}
