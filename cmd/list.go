package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/getantibody/antibody/antibody"
	"github.com/getantibody/antibody/project"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all downloaded plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := antibody.Home()
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
