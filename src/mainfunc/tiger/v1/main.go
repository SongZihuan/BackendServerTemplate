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
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
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

	err = commandlineargs.InitCommandLineArgsParser(nil)
	if err != nil {
		if errors.Is(err, commandlineargs.StopRun) {
			return exitutils.SuccessExit("server not run")
		}

		return exitutils.InitFailedError("Command Line Args Parser", err.Error())
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: commandlineargs.ConfigFile(),
		Provider:       configparser.NewYamlProvider(),
	})
	if err != nil {
		return exitutils.InitFailedError("Config file read and parser", err.Error())
	}

	if commandlineargs.OutputConfigFile() != "" {
		err = config.OutputConfig(commandlineargs.OutputConfigFile())
		if err != nil {
			return exitutils.InitFailedError("Config file output", err.Error())
		}
	}

	sigchan := signalwatcher.NewSignalExitChannel()
	defer close(sigchan)

	ser, _, err := example1.NewServerExample1(nil)
	if err != nil {
		return exitutils.InitFailedError("Server Example1", err.Error())
	}

	logger.Infof("Start to run server controller")
	go ser.Run()

	select {
	case <-sigchan:
		logger.Infof("stop by signal")
		err = nil
	case <-ser.GetCtx().Listen():
		err = ser.GetCtx().Error()
		if err == nil || errors.Is(err, servercontext.StopAllTask) {
			err = nil
			logger.Infof("stop by controller")
		} else {
			logger.Errorf("stop by controller with error")
		}
	}

	ser.Stop()

	if err != nil {
		return exitutils.RunError(err.Error())
	}

	return exitutils.SuccessExit("all tasks are completed and the main go routine exits")
}
