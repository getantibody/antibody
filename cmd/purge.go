package cmd

import (
	"fmt"
	"os"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/getantibody/antibody/project"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:     "purge",
	Aliases: []string{"rm"},
	Short:   "removes a dependency from your filesystem",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Removing", args[0])
		return os.RemoveAll(project.New(antibodylib.Home(), args[0]).Folder())
	},
}

func init() {
	RootCmd.AddCommand(purgeCmd)
}
