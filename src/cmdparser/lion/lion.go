// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lion

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/mainfunc/lion"
	"github.com/SongZihuan/BackendServerTemplate/src/mainfunc/restart"
	restartinfo "github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
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
		Short:         "Multi-tasking background system",
		Long:          "A multi-task background system controlled by a controller to run multiple tasks concurrently",
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

			return lion.Main(cmd, args, inputConfigFilePath, ppid)
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
