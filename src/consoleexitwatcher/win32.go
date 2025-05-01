// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

package consoleexitwatcher

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"syscall"
	"time"
)

var consoleexitchan = make(chan any)
var consolewaitexitchan = make(chan any)
var lastEvent consoleutils.Event = nil
var handler uintptr

func init() {
	handler = syscall.NewCallback(func(event uint) uintptr {
		if getHandler(event) {
			return 1
		}
		return 0
	})
}

func NewWin32ConsoleExitChannel() (chan any, chan any, error) {
	if !config.Data().Win32Console.Use || !consoleutils.HasConsoleWindow() {
		return consoleexitchan, consolewaitexitchan, nil
	}

	err := consoleutils.SetConsoleCtrlHandlerByCallback(handler, true)
	if err != nil {
		return nil, nil, err
	}

	return consoleexitchan, consolewaitexitchan, nil
}

func GetExitEvent() consoleutils.Event {
	if lastEvent == nil {
		logger.Panicf("not event")
	}
	return lastEvent
}

// 控制台事件处理函数（回调）
func getHandler(event uint) (res bool) {
	if lastEvent != nil {
		_ = consoleutils.SetConsoleCtrlHandlerByCallback(handler, false)
		return false
	}

	defer func() {
		if res {
			_ = consoleutils.SetConsoleCtrlHandlerByCallback(handler, false)
		}
	}()

	switch event {
	case consoleutils.CTRL_CLOSE_EVENT.GetCode(), consoleutils.CTRL_LOGOFF_EVENT.GetCode(), consoleutils.CTRL_SHUTDOWN_EVENT.GetCode():
		lastEvent = consoleutils.EventMap[event]
		close(consoleexitchan)

		if config.Data().Win32Console.ConsoleCloseRecovery.IsEnable(false) {
			err := consoleutils.MakeNewConsole(consoleutils.CodePageUTF8)
			if err != nil {
				logger.Errorf("win32 make new console failed: %s", err.Error())
			}

			logger.Warnf("The terminal will be restarted temporarily, waiting for the program to be cleaned up. Please do not close the current terminal!")
			logger.Warnf("If you do not want to restart the terminal, you can turn it off in the configuration file.")
		}
	case consoleutils.CTRL_C_EVENT.GetCode():
		if config.Data().Win32Console.CtrlCExit.IsEnable(true) {
			lastEvent = consoleutils.CTRL_C_EVENT
			close(consoleexitchan)
		}
	case consoleutils.CTRL_BREAK_EVENT.GetCode():
		if config.Data().Win32Console.CtrlBreakExit.IsEnable(true) {
			lastEvent = consoleutils.CTRL_BREAK_EVENT
			close(consoleexitchan)
		}
	default:
		logger.Errorf("unknown event: %d\n", event)
		return false
	}

	select {
	case <-consolewaitexitchan:
		logger.Warnf("Windows Console - Exit cleanup complete")
	case <-time.After(4500 * time.Millisecond):
		logger.Errorf("Windows Console - Exit cleanup timeout... (%s)", strconvutils.TimeDurationToString(4500*time.Millisecond))
	}
	return true
}
