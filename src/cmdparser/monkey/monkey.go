// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monkey

import (
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/mainfunc/monkey"
	"github.com/spf13/cobra"
)

const (
	args1Info       = "info"
	args1Install    = "install"
	args1Uninstall1 = "uninstall"
	args1Uninstall2 = "remove"
	args1Start      = "start"
	args1Stop       = "stop"
	args1Restart    = "restart"
)

var inputConfigFilePath string = "config.yaml"

func GetMainCommand() *cobra.Command {
	cmd := GetCommand(rtdata.GetName())
	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD)
	return cmd
}

func GetCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           name,
		Short:         "System service registration tool",
		Long:          "Register this software as a system service, mainly used in Windows",
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
			return monkey.Main(cmd, args, inputConfigFilePath)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.Flags().StringVarP(&inputConfigFilePath, "config", "c", inputConfigFilePath, "the file path of the configuration file")

	info := &cobra.Command{
		Use:           args1Info,
		Short:         "Show the info about the service",
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
			return monkey.MainInfo(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	install := &cobra.Command{
		Use:                args1Install,
		Short:              "Install/Register the service",
		Long:               "",
		SilenceUsage:       false,
		SilenceErrors:      false,
		DisableFlagParsing: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return monkey.MainInstall(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

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
			return monkey.MainUnInstall(cmd, args)
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
			return monkey.MainStart(cmd, args)
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
			return monkey.MainStop(cmd, args)
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
			return monkey.MainRestart(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.AddCommand(info, install, uninstall, start, stop, restart)
	return cmd
}
