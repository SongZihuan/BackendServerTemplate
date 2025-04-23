// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// 必须明确导入 prerun 包 （虽然下面的import确实导入了 prerun 包，但此处重复写一次表示冗余，以免在某些情况下本包不适用 prerun 后，下方的导入被自动删除）
import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/cmd/prerun"
)

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/prerun"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	catv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/cat/v1"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

const (
	args1Install    = "install"
	args1Uninstall1 = "uninstall"
	args1Uninstall2 = "remove"
	args1Start      = "start"
	args1Stop       = "stop"
	args1Restart    = "restart"
)

var name string = global.Name
var inputConfigFilePath string = "config.yaml"

func main() {
	command().Init().Exit()
}

func command() exitutils.ExitCode {
	err := prerun.PreRun()
	defer prerun.PostRun() // 此处defer放在err之前（因为RPreRun包含启动东西太多，虽然返回err，但不代表全部操作没成功，因此defer设置在这个位置，确保清理函数被调用。清理函数可以判断当前是否需要清理）
	if err != nil {
		return exitutils.ErrorToExit(err)
	}

	cmd := &cobra.Command{
		Use:           global.Name,
		Short:         "System service registration tool",
		Long:          "Register this software as a system service, mainly used in Windows",
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1(cmd, args, inputConfigFilePath)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.Flags().StringVarP(&inputConfigFilePath, "config", "c", inputConfigFilePath, "the file path of the configuration file")

	install := &cobra.Command{
		Use:           args1Install,
		Short:         "Install/Register the service",
		Long:          "",
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1Install(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}
	install.FParseErrWhitelist.UnknownFlags = true

	uninstall := &cobra.Command{
		Use:           args1Uninstall1,
		Short:         "Uninstall/Remove the service",
		Long:          "",
		Aliases:       []string{args1Uninstall2},
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1UnInstall(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	start := &cobra.Command{
		Use:           args1Start,
		Short:         "Start the service",
		Long:          "",
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1Start(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	stop := &cobra.Command{
		Use:           args1Stop,
		Short:         "Stop the service",
		Long:          "",
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1Stop(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	restart := &cobra.Command{
		Use:           args1Restart,
		Short:         "Restart the service",
		Long:          "",
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return catv1.MainV1Restart(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD, install, uninstall, start, stop, restart)
	return exitutils.ExitQuite(cmd.Execute())
}
