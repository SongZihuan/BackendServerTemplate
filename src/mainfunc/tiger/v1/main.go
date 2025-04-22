// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/consolewatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

var InputConfigFilePath string = "config.yaml"
var OutputConfigFilePath string = ""
var AutoReload bool = false

func MainV1(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	configProvider, err := configparser.NewProvider(InputConfigFilePath, nil)
	if err != nil {
		return exitutils.InitFailedError("Get config file provider", err.Error())
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: InputConfigFilePath,
		OutputFilePath: OutputConfigFilePath,
		Provider:       configProvider,
	})
	if err != nil {
		return exitutils.InitFailedError("Config file read and parser", err.Error())
	}

	sigchan := signalwatcher.NewSignalExitChannel()

	consolechan, consolewaitexitchan, err := consolewatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailedError("Win32 console channel", err.Error())
	}

	ser, _, err := example1.NewServerExample1(&example1.ServerExample1Option{
		StopWaitTime: config.Data().Server.StopWaitTimeDuration,
	})
	if err != nil {
		return exitutils.InitFailedError("Server Example1", err.Error())
	}

	logger.Infof("Start to run server example 1")
	go ser.Run()

	var stopErr error

SELECT:
	select {
	case <-restart.RestartChan:
		if AutoReload {
			logger.Warnf("stop/restart by config file change")
			err = nil
			stopErr = nil
		} else {
			goto SELECT
		}
	case sig := <-sigchan:
		logger.Warnf("stop by signal (%s)", sig.String())
		err = nil
		stopErr = nil
	case event := <-consolechan:
		logger.Infof("stop by console event (%s)", event.String())
		err = nil
		stopErr = nil
	case <-ser.GetCtx().Listen():
		err = ser.GetCtx().Error()
		if err == nil || errors.Is(err, servercontext.StopAllTask) {
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

	return exitutils.SuccessExit("all tasks are completed and the main go routine exits")
}
