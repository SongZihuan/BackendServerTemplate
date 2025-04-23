// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// 必须明确导入 prerun 包 （虽然下面的import确实导入了 prerun 包，但此处重复写一次表示冗余，以免在某些情况下本包不适用 prerun 后，下方的导入被自动删除）
import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/cmd/prerun"
)

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/prerun"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	lionv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/lion/v1"
	restartv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/restart/v1"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

var inputConfigFilePath string = "config.yaml"
var name string = global.Name
var reload bool = false
var ppid int = 0

func main() {
	err := prerun.PreRun()
	if err != nil {
		exitutils.Exit(err)
	}
	defer prerun.PostRun()

	cmd := &cobra.Command{
		Use:           global.Name,
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

			if reload && cmd.Flags().Changed(restart.RestartFlag) {
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
				return restartv1.MainV1(cmd, args, inputConfigFilePath)
			}

			return lionv1.MainV1(cmd, args, inputConfigFilePath, ppid)
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

	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD)

	cmd.PersistentFlags().StringVarP(&name, "name", "n", global.Name, "the program display name")

	cmd.Flags().BoolVar(&reload, "auto-reload", false, "auto reload config file when the file changed")
	cmd.Flags().IntVar(&ppid, restart.RestartFlag, 0, "restart mode, note: DO NOT SET THIS FLAG unless you know your purpose clearly.")

	cmd.Flags().StringVarP(&inputConfigFilePath, "config", "c", inputConfigFilePath, "the file path of the configuration file")

	exitutils.ExitQuite(cmd.Execute())
}
