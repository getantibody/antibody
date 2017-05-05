package cmd

import (
	"fmt"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "shows current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("antibody version %v\n", antibodylib.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
