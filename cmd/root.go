package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd is the cobra root command
var RootCmd = &cobra.Command{
	Use:   "antibody",
	Short: "The fastest shell plugin manager.",
	Long: `Antibody can manage plugins for shells (zsh, for example),
both loading them with source or export-ing them to PATH.`,
}

// Execute executes the main cobra command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
