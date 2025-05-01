// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tiger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/mainfunc/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/mainfunc/tiger"
	restartinfo "github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/spf13/cobra"
)

var inputConfigFilePath string = "config.yaml"
var name string = global.Name
var reload bool = false
var ppid int = 0

func GetMainCommand() *cobra.Command {
	cmd := GetCommand(global.Name)
	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD)
	return cmd
}

func GetCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           name,
		Short:         "Single-tasking background system",
		Long:          "A single-task background system that runs a single task directly without using a controller",
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false

			if name = cleanstringutils.GetStringOneLine(name); cmd.Flags().Changed("name") && name != "" {
				global.Name = name
				global.NameFlagChanged = true
			} else {
				global.NameFlagChanged = false
			}

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false

			if reload && cmd.Flags().Changed(restartinfo.RestartFlag) {
				if ppid == 0 {
					return fmt.Errorf("`restart` cannot be specified as 0")
				}
				reload = false
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			if reload {
				return restart.Main(cmd, args, inputConfigFilePath)
			}

			return tiger.Main(cmd, args, inputConfigFilePath, ppid)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&name, "name", "n", global.Name, "the program display name")

	cmd.Flags().BoolVar(&reload, "auto-reload", false, "auto reload config file when the file changed")
	cmd.Flags().IntVar(&ppid, restartinfo.RestartFlag, 0, "restart mode, note: DO NOT SET THIS FLAG unless you know your purpose clearly.")

	cmd.Flags().StringVarP(&inputConfigFilePath, "config", "c", inputConfigFilePath, "the file path of the configuration file")

	return cmd
}
