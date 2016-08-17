package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody"
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
	var input []string
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(ctx.Args()) == 0 {
		entries, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		input = strings.Split(string(entries), "\n")
	} else {
		input = ctx.Args()
	}
	sh, err := antibody.New(gohome.Cache("antibody")+"/", input).Bundle()
	if err != nil {
		return err
	}
	fmt.Println(sh)
	return nil
}
