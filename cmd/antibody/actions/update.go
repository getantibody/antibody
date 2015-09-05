package actions

import (
	"github.com/caarlos0/antibody"
	"github.com/codegangsta/cli"
)

// Update all installed bundles
func Update(c *cli.Context) {
	antibody.New(installedBundles()).Update()
}
