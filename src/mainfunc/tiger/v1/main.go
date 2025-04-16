// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/commandlineargs"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/consolewatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/server/example1"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func MainV1() (exitCode exitutils.ExitCode) {
	var err error

	err = consoleutils.SetConsoleCPSafe(consoleutils.CodePageUTF8)
	if err != nil {
		return exitutils.InitFailedErrorForWin32ConsoleModule(err.Error())
	}

	err = logger.InitBaseLogger(loglevel.LevelDebug, true, nil, nil, nil, nil)
	if err != nil {
		return exitutils.InitFailedErrorForLoggerModule(err.Error())
	}
	defer logger.CloseLogger()
	defer logger.Recover()

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
	select {
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
