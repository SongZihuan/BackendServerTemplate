// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tiger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/consoleexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/sigexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string, inputConfigFilePath string, ppid int) (exitCode error) {
	var err error

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: inputConfigFilePath,
		AutoReload:     ppid != 0,
	})
	if err != nil {
		return exitutils.InitFailed("Config file read and parser", err.Error())
	}

	sigexitchan := sigexitwatcher.GetSignalExitChannelFromConfig()

	consoleexitchan, consolewaitexitchan, err := consoleexitwatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailed("Win32 console channel", err.Error())
	}

	ppidchan := restart.PpidWatcher(ppid)

	ser, _, err := example1.NewServerExample1(&example1.ServerExample1Option{
		StopWaitTime: config.Data().Server.StopWaitTimeDuration,
	})
	if err != nil {
		return exitutils.InitFailed("Server Example1", err.Error())
	}

	logger.Infof("Start to run server example 1")
	go ser.Run()

	var stopErr error

	select {
	case <-restart.RestartChan:
		if ppid != 0 {
			logger.Warnf("stop to restart")
			err = nil
			stopErr = nil
		} else {
			logger.Warnf("stop to restart (error: restart not set)")
			err = fmt.Errorf("stop by restart, but restart not set")
			stopErr = err
		}
	case <-ppidchan:
		if ppid != 0 {
			logger.Warnf("stop by parent process")
			err = nil
			stopErr = nil
		} else {
			logger.Warnf("stop by parent process (error: ppid not set)")
			err = fmt.Errorf("stop by parent process, but pppid not set")
			stopErr = err
		}
	case <-sigexitchan:
		logger.Warnf("stop by signal (%s)", sigexitwatcher.GetExitSignal().String())
		err = nil
		stopErr = nil
	case <-consoleexitchan:
		logger.Infof("stop by console event (%s)", consoleexitwatcher.GetExitEvent().String())
		err = nil
		stopErr = nil
	case <-ser.GetCtx().Listen():
		err = ser.GetCtx().Error()
		if err == nil {
			logger.Infof("stop by server")
			err = nil
			stopErr = nil
		} else {
			logger.Errorf("stop by server with error: %s", err.Error())
			stopErr = err
		}
	}

	ser.Stop()
	close(consolewaitexitchan)

	if stopErr != nil {
		return exitutils.RunError(stopErr.Error())
	}

	select {
	case <-restart.RestartChan:
		return exitutils.SuccessExit("all tasks are completed and the main go routine exits", exitutils.ExitCodeReload)
	default:
		return exitutils.SuccessExit("all tasks are completed and the main go routine exits")
	}
}
