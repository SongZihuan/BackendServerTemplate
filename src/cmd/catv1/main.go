// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/globalmain"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/subcmd"
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	catv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/cat/v1"
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

func main() {
	defer logger.Recover()

	cmd := root.GetRootCMD("System service registration tool",
		"Register this software as a system service, mainly used in Windows.",
		nil,
		false,
		catv1.MainV1)

	subcmd.AddSubCMDOfRoot(cmd)
	cmd.Flags().StringVarP(&catv1.InputConfigFilePath, "config", "c", catv1.InputConfigFilePath, "the file path of the configuration file")
	cmd.Flags().StringVarP(&catv1.OutputConfigFilePath, "output-config", "o", catv1.OutputConfigFilePath, "the file path of the output configuration file")

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

			err := globalmain.PreRun(false)
			if err != nil {
				return err
			}

			return catv1.MainV1Install(cmd, args)
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

			err := globalmain.PreRun(false)
			if err != nil {
				return err
			}

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

			err := globalmain.PreRun(false)
			if err != nil {
				return err
			}

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

			err := globalmain.PreRun(false)
			if err != nil {
				return err
			}

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

			err := globalmain.PreRun(false)
			if err != nil {
				return err
			}

			return catv1.MainV1Restart(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.AddCommand(install, uninstall, start, stop, restart)
	exitutils.ExitQuite(cmd.Execute())
}
