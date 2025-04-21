// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/globalmain"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/consolewatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/controller"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example2"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

var InputConfigFilePath string = "config.yaml"
var OutputConfigFilePath string = ""

func MainV1(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = globalmain.PreRun()
	if err != nil {
		return err
	}
	defer globalmain.PostRun()

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: InputConfigFilePath,
		OutputFilePath: OutputConfigFilePath,
		Provider:       configparser.NewYamlProvider(),
	})
	if err != nil {
		return exitutils.InitFailedError("Config file read and parser", err.Error())
	}

	sigchan := signalwatcher.NewSignalExitChannel()

	consolechan, consolewaitexitchan, err := consolewatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailedError("Win32 console channel", err.Error())
	}

	ctrl, err := controller.NewController(&controller.ControllerOption{
		StopWaitTime: config.Data().Server.StopWaitTimeDuration,
	})
	if err != nil {
		return exitutils.InitFailedError("Server Controller", err.Error())
	}

	ser1, _, err := example1.NewServerExample1(nil)
	if err != nil {
		return exitutils.InitFailedError("Server Example1", err.Error())
	}

	err = ctrl.AddServer(ser1)
	if err != nil {
		return exitutils.InitFailedError("Add Server Example1", err.Error())
	}

	ser2, _, err := example2.NewServerExample2(nil)
	if err != nil {
		return exitutils.InitFailedError("Server Example2", err.Error())
	}

	err = ctrl.AddServer(ser2)
	if err != nil {
		return exitutils.InitFailedError("Add Server Example2", err.Error())
	}

	logger.Infof("Start to run server controller")
	go ctrl.Run()

	var stopErr error
	select {
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

	return exitutils.SuccessExit("all tasks are completed and the main go routine exits")
}
