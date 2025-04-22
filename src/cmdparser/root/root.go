// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package root

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/globalmain"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/spf13/cobra"
)

const (
	ConsoleModeNormal = "normal"
	ConsoleModeNO     = "no"
)

var name string = global.Name
var nameChanged bool = false
var isRestart bool = false
var consoleMode string = ConsoleModeNormal
var hasConsole bool = false

func GetRootCMD(shortDescribe string, longDescribe string, reload *bool, _hasConsole bool, action func(cmd *cobra.Command, args []string) error) *cobra.Command {
	hasConsole = _hasConsole && consoleutils.HasConsoleWindow()

	cmd := &cobra.Command{
		Use:           global.Name,
		Short:         shortDescribe,
		Long:          longDescribe,
		SilenceUsage:  false,
		SilenceErrors: false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false

			if name = cleanstringutils.GetStringOneLine(name); cmd.Flags().Changed("name") && name != "" {
				global.Name = name
				nameChanged = true
			} else {
				nameChanged = false
			}

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false

			if hasConsole {
				if consoleMode != ConsoleModeNormal && consoleMode != ConsoleModeNO {
					return fmt.Errorf("console-mode must be %s or %s", ConsoleModeNormal, ConsoleModeNO)
				}
			} else if _hasConsole && consoleMode != ConsoleModeNO {
				return fmt.Errorf("console-mode must be %s, there is not console be found", ConsoleModeNO)
			}

			if reload != nil && *reload {
				if consoleMode == ConsoleModeNO {
					return fmt.Errorf("`auto-reload` can only be enabled when `console-mode` is `normal`")
				}

				if cmd.Flags().Changed(restart.RestartFlag) {
					if isRestart {
						err := restart.FromRestart()
						if err != nil {
							return fmt.Errorf("restart failed to attach console: %s", err)
						}
					} else {
						return fmt.Errorf("`restart` cannot be specified as false")
					}
				} else {
					return restart.FirstRun()
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			err := globalmain.PreRun(HasConsole())
			if err != nil {
				return err
			}

			return action(cmd, args)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			globalmain.PostRun()
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = false
			cmd.SilenceErrors = false
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&name, "name", "n", global.Name, "the program display name")

	if _hasConsole {
		cmd.Flags().StringVar(&consoleMode, "console-mode", ConsoleModeNormal, "the console mode. normally, select `normal`. If you are not using a terminal (e.g. redirection, etc.) please select `no`.")
	}

	if reload != nil {
		cmd.Flags().BoolVar(reload, "auto-reload", false, "auto reload config file when the file changed")
		cmd.Flags().BoolVar(&isRestart, restart.RestartFlag, false, "restart mode, note: DO NOT SET THIS FLAG unless you know your purpose clearly.")
	}

	return cmd
}

func Name() string {
	return name
}

func NameChanged() bool {
	return nameChanged
}

func IsRestart() bool {
	return isRestart
}

func ConsoleMode() string {
	switch consoleMode {
	case ConsoleModeNormal:
		return ConsoleModeNormal
	case ConsoleModeNO:
		return ConsoleModeNO
	default:
		return ConsoleModeNormal
	}
}

func HasConsole() bool {
	return hasConsole && ConsoleMode() == ConsoleModeNormal
}
