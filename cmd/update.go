package cmd

import (
	"fmt"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/getantibody/antibody/project"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "updates all previously bundled plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		var home = antibodylib.Home()
		fmt.Printf("Updating all bundles in %v...\n", home)
		return project.Update(home)
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
