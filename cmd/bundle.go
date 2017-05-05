package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var bundleCmd = &cobra.Command{
	Use:     "bundle",
	Aliases: []string{"b"},
	Short:   "downloads a bundle and prints its source line",
	RunE: func(cmd *cobra.Command, args []string) error {
		var input io.Reader
		if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(args) == 0 {
			input = os.Stdin
		} else {
			input = bytes.NewBufferString(strings.Join(args, " "))
		}
		sh, err := antibodylib.New(antibodylib.Home(), input).Bundle()
		if err != nil {
			return err
		}
		fmt.Println(sh)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(bundleCmd)
}
