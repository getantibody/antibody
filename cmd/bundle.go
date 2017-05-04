// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/getantibody/antibody/antibody"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "downloads (if needed) and then prints the needed shell commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		var input io.Reader
		if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(args) == 0 {
			input = os.Stdin
		} else {
			input = bytes.NewBufferString(strings.Join(args, " "))
		}
		sh, err := antibody.New(antibody.Home(), input).Bundle()
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
