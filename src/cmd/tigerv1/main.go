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
	restartv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/restart/v1"
	tigerv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/tiger/v1"
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
	command().ClampAttribute().Exit()
}

func command() exitutils.ExitCode {
	err := prerun.PreRun()
	defer prerun.PostRun() // 此处defer放在err之前（因为RPreRun包含启动东西太多，虽然返回err，但不代表全部操作没成功，因此defer设置在这个位置，确保清理函数被调用。清理函数可以判断当前是否需要清理）
	if err != nil {
		return exitutils.ErrorToExit(err)
	}

	cmd := &cobra.Command{
		Use:           global.Name,
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

			return tigerv1.MainV1(cmd, args, inputConfigFilePath, ppid)
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

	return exitutils.ErrorToExitQuite(cmd.Execute())
}
