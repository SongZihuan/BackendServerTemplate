// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

package consolewatcher

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"time"
)

// 控制台事件处理函数（回调）
func consoleHandler(exitChannel chan consoleutils.Event, waitExitChannel chan any) func(event uint) bool {
	return func(event uint) bool {
		switch event {
		case consoleutils.CTRL_CLOSE_EVENT.GetCode(), consoleutils.CTRL_LOGOFF_EVENT.GetCode(), consoleutils.CTRL_SHUTDOWN_EVENT.GetCode():
			exitChannel <- consoleutils.EventMap[event]

			if config.Data().Win32Console.ConsoleCloseRecovery.IsEnable(false) {
				err := consoleutils.MakeNewConsole(consoleutils.CodePageUTF8)
				if err != nil {
					logger.Errorf("win32 make new console failed: %s", err.Error())
				}
			}

			logger.Warnf("终端暂时重启，等待程序清理完毕，请勿关闭当前终端！")
			logger.Warnf("若不希望重启终端，可在配置文件处关闭。")

			select {
			case <-waitExitChannel:
				// pass
			case <-time.After(4500 * time.Millisecond):
				logger.Errorf("Windows Console - 退出清理超时... (%s)", strconvutils.TimeDurationToString(4500*time.Millisecond))
			}
			return true
		case consoleutils.CTRL_C_EVENT.GetCode():
			if config.Data().Win32Console.CtrlCExit.IsEnable(true) {
				exitChannel <- consoleutils.CTRL_C_EVENT
			}
			return true
		case consoleutils.CTRL_BREAK_EVENT.GetCode():
			if config.Data().Win32Console.CtrlBreakExit.IsEnable(true) {
				exitChannel <- consoleutils.CTRL_BREAK_EVENT
			}
			return true
		default:
			logger.Errorf("未知事件: %d\n", event)
			return false
		}
	}
}

func NewWin32ConsoleExitChannel() (chan consoleutils.Event, chan any, error) {
	var exitChannel = make(chan consoleutils.Event)
	var waitExitChannel = make(chan any)

	if !config.Data().Win32Console.Use {
		return exitChannel, waitExitChannel, nil
	}

	err := consoleutils.SetConsoleCtrlHandler(consoleHandler(exitChannel, waitExitChannel), true)
	if err != nil {
		return nil, nil, err
	}

	return exitChannel, waitExitChannel, nil
}
