package antibody

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func Clone(bundle string, home string) (string, error) {
	folder := home + strings.Replace(bundle, "/", "-", -1)
	var cloneErr error
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--depth", "1", GH+bundle, folder)
		cloneErr = clone.Run()
	}
	return folder, cloneErr
}

func Pull(bundle string, home string) string {
	folder := home + bundle
	pull := exec.Command("git", "-C", folder, "pull", "origin", "master")
	pull.Run()
	return folder
}

func Update(home string) {
	bundles, _ := ioutil.ReadDir(home)
	for _, bundle := range bundles {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			fmt.Println(Pull(bundle.Name(), home))
		}
	}
}
