package antibody

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var version = "master"

func bundle(bundle string, home string) {
	NewAntibody([]Bundle{NewGitBundle(bundle, home)}).Download()
}

// ProcessStdin processes the OS SDTDIN.
func ProcessStdin(stdin io.Reader, home string) {
	entries, _ := ioutil.ReadAll(stdin)
	var bundles []Bundle
	for _, bundle := range strings.Split(string(entries), "\n") {
		if bundle == "" {
			continue
		}
		bundles = append(bundles, NewGitBundle(bundle, home))
	}
	NewAntibody(bundles).Download()
}

func update(home string) {
	entries, _ := ioutil.ReadDir(home)
	var bundles []Bundle
	for _, bundle := range entries {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			bundles = append(bundles, NewGitBundle(bundle.Name(), home))
		}
	}
	NewAntibody(bundles).Update()
}

// ProcessArgs processes arguments passed to the executable.
func ProcessArgs(args []string, home string) {
	cmd := args[0]
	if cmd == "update" {
		update(home)
	} else if cmd == "bundle" {
		bundle(args[1], home)
	} else if cmd == "version" {
		fmt.Println(version)
	} else {
		panic("Invalid command: " + cmd)
	}
}

// ReadStdin checks if there is something being passed to the STDIN
func ReadStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// Home returns the ANTIBODY_HOME to use, wether it is the default or another
// one.
func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = filepath.Join(os.Getenv("HOME"), ".antibody")
	}
	return home
}
