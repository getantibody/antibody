package cmd

import (
	"fmt"

	"github.com/getantibody/antibody/antibody"
	"github.com/spf13/cobra"
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:     "home",
	Short:   "shows the current antibody home folder",
	Aliases: []string{"prefix", "p"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(antibody.Home())
	},
}

func init() {
	RootCmd.AddCommand(homeCmd)
}
