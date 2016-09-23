package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

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
	var input io.Reader
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(ctx.Args()) == 0 {
		input = os.Stdin
	} else {
		input = bytes.NewBufferString(strings.Join(ctx.Args(), "\n"))
	}
	sh, err := antibody.New(antibody.Home(), input).Bundle()
	if err != nil {
		return err
	}
	fmt.Println(sh)
	return nil
}
