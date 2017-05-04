package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/getantibody/antibody/project"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list all downloaded plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := antibodylib.Home()
		projects, err := project.List(home)
		if err != nil {
			return err
		}
		for _, b := range projects {
			fmt.Println(filepath.Join(home, b))
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
