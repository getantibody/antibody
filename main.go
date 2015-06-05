package antibody

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
		clone.Run()
	}
	return folder
}

func pull(bundle string, home string) string {
	folder := home + bundle
	pull := exec.Command("git", "-C", folder, "pull", "origin", "master")
	pull.Run()
	return folder
}

func update(home string) {
	bundles, _ := ioutil.ReadDir(home)
	for _, bundle := range bundles {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			fmt.Println(pull(bundle.Name(), home))
		}
	}
}

func processStdin(home string) {
	bundles, _ := ioutil.ReadAll(os.Stdin)
	for _, bundle := range strings.Split(string(bundles), "\n") {
		if bundle != "" {
			fmt.Println(clone(bundle, home))
		}
	}
}

func processArgs(home string) {
	cmd := os.Args[1:][0]
	if cmd == "update" {
		update(home)
	} else if cmd == "bundle" {
		fmt.Println(clone(os.Args[1:][1], home))
	} else {
		panic("Invalid command: " + cmd)
	}
}

func readStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = os.Getenv("HOME") + "/.antibody/"
	}
	return home
}

func main() {
	home := Home()
	if readStdin() {
		processStdin(home)
	} else {
		processArgs(home)
	}
}
