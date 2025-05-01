// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package restart

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/consoleexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/sigexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/utils/exitutils"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string, inputConfigFilePath string) (exitCode error) {
	var err error

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: inputConfigFilePath,
		AutoReload:     false,
	})
	if err != nil {
		return exitutils.InitFailed("Config file read and parser", err.Error())
	}

	sigexitchan := sigexitwatcher.GetSignalExitChannelFromConfig()

	consoleexitchan, consolewaitexitchan, err := consoleexitwatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailed("Win32 console channel", err.Error())
	}

	stopchan := restart.RunRestart()

	var stopErr error

	select {
	case <-stopchan:
		logger.Warnf("stop by sub process")
		err = nil
		stopErr = nil
	case <-sigexitchan:
		logger.Warnf("stop by signal (%s)", sigexitwatcher.GetExitSignal().String())
		err = nil
		stopErr = nil
	case <-consoleexitchan:
		logger.Infof("stop by console event (%s)", consoleexitwatcher.GetExitEvent().String())
		err = nil
		stopErr = nil
	}

	close(consolewaitexitchan)

	if stopErr != nil {
		return exitutils.RunError(stopErr.Error())
	}

	return exitutils.SuccessExit("all tasks are completed")
}
