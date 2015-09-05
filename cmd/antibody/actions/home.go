package actions

import (
	"fmt"

	"github.com/caarlos0/antibody"
	"github.com/codegangsta/cli"
)

// Home shows current antibody home folder
func Home(ctx *cli.Context) {
	fmt.Println(antibody.Home())
}
