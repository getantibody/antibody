package antibody

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Bundle(bundle string, home string) {
	folder, err := Clone(bundle, home)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(folder)
}

func ProcessStdin(home string) {
	bundles, _ := ioutil.ReadAll(os.Stdin)
	for _, bundle := range strings.Split(string(bundles), "\n") {
		if bundle != "" {
			go Bundle(bundle, home)
		}
	}
}

func ProcessArgs(args []string, home string) {
	cmd := args[0]
	if cmd == "update" {
		go Update(home)
	} else if cmd == "bundle" {
		go Bundle(args[1], home)
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
	}
	return home
}
