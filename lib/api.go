package antibody

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func DoBundle(bundle string, home string) {
	NewAntibody([]Bundle{NewGitBundle(bundle, home)}).Download()
}

func ProcessStdin(stdin io.Reader, home string) {
	entries, _ := ioutil.ReadAll(stdin)
	bundles := make([]Bundle, 0)
	for _, bundle := range strings.Split(string(entries), "\n") {
		if bundle == "" {
			continue
		}
		bundles = append(bundles, NewGitBundle(bundle, home))
	}
	NewAntibody(bundles).Download()
}

func Update(home string) {
	entries, _ := ioutil.ReadDir(home)
	var bundles []Bundle
	for _, bundle := range entries {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			bundles = append(bundles, NewGitBundle(bundle.Name(), home))
		}
	}
	NewAntibody(bundles).Update()
}

func ProcessArgs(args []string, home string) {
	cmd := args[0]
	if cmd == "update" {
		Update(home)
	} else if cmd == "bundle" {
		DoBundle(args[1], home)
	} else {
		panic("Invalid command: " + cmd)
	}
}

func ReadStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = os.Getenv("HOME") + "/.antibody/"
	} else {
		if !strings.HasSuffix(home, "/") {
			home += "/"
		}
	}
	return home
}
