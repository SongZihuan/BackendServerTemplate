// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/consolewatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/server/controller"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example2"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

func MainV1(cmd *cobra.Command, args []string, inputConfigFilePath string, ppid int) (exitCode error) {
	var err error

	configProvider, err := configparser.NewProvider(inputConfigFilePath, &configparser.NewProviderOption{
		AutoReload: ppid != 0,
	})
	if err != nil {
		return exitutils.InitFailed("Get config file provider", err.Error())
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: inputConfigFilePath,
		Provider:       configProvider,
	})
	if err != nil {
		return exitutils.InitFailed("Config file read and parser", err.Error())
	}

	sigchan := signalwatcher.NewSignalExitChannel()

	consolechan, consolewaitexitchan, err := consolewatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailed("Win32 console channel", err.Error())
	}

	ppidchan := restart.PpidWatcher(ppid)

	ctrl, err := controller.NewController(&controller.ControllerOption{
		StopWaitTime: config.Data().Server.StopWaitTimeDuration,
	})
	if err != nil {
		return exitutils.InitFailed("Server Controller", err.Error())
	}

	ser1, _, err := example1.NewServerExample1(&example1.ServerExample1Option{
		LockThread: true,
	})
	if err != nil {
		return exitutils.InitFailed("Server Example1", err.Error())
	}

	err = ctrl.AddServer(ser1)
	if err != nil {
		return exitutils.InitFailed("Add Server Example1", err.Error())
	}

	ser2, _, err := example2.NewServerExample2(nil)
	if err != nil {
		return exitutils.InitFailed("Server Example2", err.Error())
	}

	err = ctrl.AddServer(ser2)
	if err != nil {
		return exitutils.InitFailed("Add Server Example2", err.Error())
	}

	logger.Infof("Start to run server controller")
	go ctrl.Run()

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
	case sig := <-sigchan:
		logger.Warnf("stop by signal (%s)", sig.String())
		err = nil
		stopErr = nil
	case event := <-consolechan:
		logger.Infof("stop by console event (%s)", event.String())
		err = nil
		stopErr = nil
	case <-ctrl.GetCtx().Listen():
		err = ctrl.GetCtx().Error()
		if err == nil || errors.Is(err, servercontext.StopAllTask) {
			logger.Infof("stop by controller")
			err = nil
			stopErr = nil
		} else {
			logger.Errorf("stop by controller with error: %s", err.Error())
			stopErr = err
		}
	}

	ctrl.Stop()
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
