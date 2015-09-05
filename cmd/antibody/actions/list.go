package actions

import (
	"fmt"
	"io/ioutil"

	"github.com/caarlos0/antibody"
	"github.com/codegangsta/cli"
)

// List all installed bundles
func List(c *cli.Context) {
	for _, bundle := range installedBundles() {
		fmt.Println(bundle.Name())
	}
}

func installedBundles() []antibody.Bundle {
	home := antibody.Home()
	entries, _ := ioutil.ReadDir(home)
	var bundles []antibody.Bundle
	for _, bundle := range entries {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			bundles = append(
				bundles,
				antibody.NewBundle(bundle.Name(), home),
			)
		}
	}
	return bundles
}
