// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/commandlineargs"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/server/controller"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example2"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func MainV1() (exitCode int) {
	var err error

	err = logger.InitBaseLogger(loglevel.LevelDebug, true, true, nil, nil)
	if err != nil {
		return exitutils.InitFailedErrorForLoggerModule(err.Error())
	}
	defer logger.CloseLogger()

	err = commandlineargs.InitCommandLineArgsParser(nil)
	if err != nil {
		if errors.Is(err, commandlineargs.StopRun) {
			return exitutils.SuccessExitQuite()
		}

		return exitutils.InitFailedError("Command Line Args Parser", err.Error())
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: commandlineargs.ConfigFile(),
		OutputFilePath: commandlineargs.OutputConfigFile(),
		Provider:       configparser.NewYamlProvider(),
	})
	if err != nil {
		return exitutils.InitFailedError("Config file read and parser", err.Error())
	}

	sigchan := signalwatcher.NewSignalExitChannel()
	defer close(sigchan)

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

	select {
	case <-sigchan:
		logger.Infof("stop by signal")
		err = nil
	case <-ctrl.GetCtx().Listen():
		err = ctrl.GetCtx().Error()
		if err == nil || errors.Is(err, servercontext.StopAllTask) {
			err = nil
			logger.Infof("stop by controller")
		} else {
			logger.Errorf("stop by controller with error")
		}
	}

	ctrl.Stop()

	if err != nil {
		return exitutils.RunError(err.Error())
	}

	return exitutils.SuccessExit("all tasks are completed and the main go routine exits")
}
