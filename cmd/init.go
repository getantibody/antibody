package cmd

import (
	"fmt"

	"github.com/getantibody/antibody/shell"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Initializes the shell so Antibody can work as expected",
	RunE: func(cmd *cobra.Command, args []string) error {
		sh, err := shell.Init()
		fmt.Println(sh)
		return err
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
