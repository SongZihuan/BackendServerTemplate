// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cat

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/consoleexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/example3"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/sigexitwatcher"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/kardianos/service"
)

type Program struct {
	sigexitchan         chan any
	consoleexitchan     chan any
	consolewaitexitchan chan any
	stopErr             error
	ser                 *server.Server
	exitCode            exitutils.ExitCode
	configPath          string
}

func NewProgram() *Program {
	return NewRunProgram("")
}

func NewRunProgram(configPath string) *Program {
	return &Program{
		sigexitchan:         make(chan any), // 临时顶替（后续会重新复制）
		consoleexitchan:     make(chan any), // 临时顶替（后续会重新复制）
		consolewaitexitchan: make(chan any), // 临时顶替（后续会重新复制）
		stopErr:             nil,
		ser:                 nil,
		exitCode:            exitutils.ExitCode(0),
		configPath:          configPath,
	}
}

func (p *Program) Start(s service.Service) error {
	var err error

	if p.configPath == "" {
		logger.Panicf("The main process should not be called.")
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: p.configPath,
		AutoReload:     false,
	})
	if err != nil {
		p.exitCode = exitutils.InitFailed("Config file read and parser", err.Error())
		return err
	}

	// 不使用 windows console, 因为注册为服务后运行是没有
	p.consoleexitchan, p.consolewaitexitchan, err = consoleexitwatcher.NewWin32ConsoleExitChannel()
	if err != nil {
		return exitutils.InitFailed("Win32 console channel", err.Error())
	}

	p.sigexitchan = sigexitwatcher.GetSignalExitChannelFromConfig()

	sercore1, err := example3.NewServerExample3Core(nil)
	if err != nil {
		return exitutils.InitFailed("Server Example1", err.Error())
	}

	p.ser, _, err = server.NewServer(&server.ServerOption{
		StopWaitTime:    config.Data().Server.Example3.StopWaitTimeDuration,
		StartupWaitTime: config.Data().Server.Example3.StartupWaitTimeDuration,
		LockThread:      config.Data().Server.Example3.LockThread.IsEnable(false),
		ServerCore:      sercore1,
	})

	logger.Infof("Start to run server %s", p.ser.Name())
	err, timeout := server.Run(p.ser)
	if err != nil {
		logger.Errorf("start server %s error: %s", p.ser.Name(), err.Error())
	} else if timeout {
		logger.Warnf("start server %s run success. but check timeout", p.ser.Name())
	} else {
		logger.Warnf("start server %s success", p.ser.Name())
	}

	go func() {
		select {
		case <-p.sigexitchan:
			logger.Warnf("Stop by signal (%s)", sigexitwatcher.GetExitSignal().String())
			err = nil
			p.stopErr = nil
		case <-p.consoleexitchan:
			logger.Infof("stop by console event (%s)", consoleexitwatcher.GetExitEvent().String())
			err = nil
			p.stopErr = nil
		case <-p.ser.GetCtx().Listen():
			err = p.ser.GetCtx().Error()
			if err == nil {
				logger.Infof("Stop by server")
				err = nil
				p.stopErr = nil
			} else {
				logger.Errorf("Stop by server with error: %s", err.Error())
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
	if p.configPath == "" {
		logger.Panicf("The main process should not be called.")
	}

	p.ser.Stop()
	if p.stopErr != nil {
		p.exitCode = exitutils.RunError(p.stopErr.Error())
		return p.stopErr
	}
	p.exitCode = exitutils.SuccessExit("all tasks are completed")
	return nil
}

func (p *Program) ExitCode() exitutils.ExitCode {
	return p.exitCode
}
