package command

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/getantibody/antibody"
	"github.com/getantibody/antibody/bundle"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Bundle downloads (if needed) and then sources a given repo
var Bundle = cli.Command{
	Name:   "bundle",
	Usage:  "downloads (if needed) and then sources a given repo",
	Action: doBundle,
}

func doBundle(ctx *cli.Context) error {
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(ctx.Args()) == 0 {
		entries, err := ioutil.ReadAll(os.Stdin)
		if err != nil || len(entries) == 0 {
			return err
		}
		antibody.New(
			bundle.Parse(string(entries), antibody.Home()),
		).Download()
	} else {
		antibody.New(
			[]bundle.Bundle{
				bundle.New(strings.Join(ctx.Args(), " "), antibody.Home()),
			},
		).Download()
	}
	return nil
}
