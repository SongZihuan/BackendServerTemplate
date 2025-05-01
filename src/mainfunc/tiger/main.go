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
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/sigexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
	"time"
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

	sercore1, err := example1.NewServerExample1Core(nil)
	if err != nil {
		return exitutils.InitFailed("Server Example1", err.Error())
	}

	ser, _, err := server.NewServer(&server.ServerOption{
		StopWaitTime:    10 * time.Second,
		StartupWaitTime: 3 * time.Second,
		ServerCore:      sercore1,
	})

	logger.Infof("Start to run server %s", ser.Name())
	err, timeout := server.Run(ser)
	if err != nil {
		logger.Errorf("start server %s error: %s", ser.Name(), err.Error())
	} else if timeout {
		logger.Warnf("start server %s run success. but check timeout", ser.Name())
	} else {
		logger.Warnf("start server %s success", ser.Name())
	}

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

	ser.StopAndWait()
	close(consolewaitexitchan)

	if stopErr != nil {
		return exitutils.RunError(stopErr.Error())
	}

	select {
	case <-restart.RestartChan:
		return exitutils.SuccessExit("restart program", exitutils.ExitCodeReload)
	default:
		return exitutils.SuccessExit("all tasks are completed")
	}
}
