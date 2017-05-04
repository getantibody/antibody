package cmd

import (
	"fmt"

	"github.com/getantibody/antibody/antibody"
	"github.com/getantibody/antibody/project"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates all previously bundled plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Updating all bundles in " + antibody.Home() + "...")
		return project.Update(antibody.Home())
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
