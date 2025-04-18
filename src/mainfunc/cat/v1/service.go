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
	"github.com/SongZihuan/BackendServerTemplate/src/server/example3"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/server/serverinterface"
	"github.com/SongZihuan/BackendServerTemplate/src/signalwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/kardianos/service"
	"os"
)

type Program struct {
	sigchan             chan os.Signal
	consolechan         chan consoleutils.Event
	consolewaitexitchan chan any
	stopErr             error
	ser                 serverinterface.Server
	exitCode            exitutils.ExitCode
}

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) Start(s service.Service) error {
	err := commandlineargs.InitCommandLineArgsParser(nil)
	if err != nil {
		if errors.Is(err, commandlineargs.StopRun) {
			p.exitCode = exitutils.SuccessExitQuite()
			return err
		}

		p.exitCode = exitutils.InitFailedError("Command Line Args Parser", err.Error())
		return err
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: commandlineargs.ConfigFile(),
		OutputFilePath: commandlineargs.OutputConfigFile(),
		Provider:       configparser.NewYamlProvider(),
	})
	if err != nil {
		p.exitCode = exitutils.InitFailedError("Config file read and parser", err.Error())
		return err
	}

	p.sigchan = signalwatcher.NewSignalExitChannel()

	p.consolechan, p.consolewaitexitchan, err = consolewatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		p.exitCode = exitutils.InitFailedError("Win32 console channel", err.Error())
		return err
	}

	p.ser, _, err = example3.NewServerExample3(&example3.ServerExample3Option{
		StopWaitTime: config.Data().Server.StopWaitTimeDuration,
	})
	if err != nil {
		return exitutils.InitFailedError("Server Example1", err.Error())
	}

	logger.Infof("Start to run server example 3")
	go p.ser.Run()
	go func() {
		select {
		case sig := <-p.sigchan:
			logger.Warnf("stop by signal (%s)", sig.String())
			err = nil
			p.stopErr = nil
		case event := <-p.consolechan:
			logger.Infof("stop by console event (%s)", event.String())
			err = nil
			p.stopErr = nil
		case <-p.ser.GetCtx().Listen():
			err = p.ser.GetCtx().Error()
			if err == nil || errors.Is(err, servercontext.StopAllTask) {
				logger.Infof("stop by server")
				err = nil
				p.stopErr = nil
			} else {
				logger.Errorf("stop by server with error: %s", err.Error())
				p.stopErr = err
			}
		}

		p.stopErr = s.Stop()
		if p.stopErr != nil {
			p.exitCode = exitutils.RunErrorQuite()
		}
	}()

	return nil
}

func (p *Program) Stop(s service.Service) error {
	p.ser.Stop()
	close(p.consolewaitexitchan)

	if p.stopErr != nil {
		p.exitCode = exitutils.RunError(p.stopErr.Error())
		return p.stopErr
	}

	p.exitCode = exitutils.SuccessExit("all tasks are completed and the main go routine exits")
	return nil
}

func (p *Program) ExitCode() exitutils.ExitCode {
	return p.exitCode
}
