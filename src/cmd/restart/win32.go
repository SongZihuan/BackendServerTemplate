// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

package restart

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/sliceutils"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var RestartChan = make(chan bool)

const RestartWaitTime = 5 * time.Second

func RestartProgram(restartFlag string) error {
	select {
	case _, ok := <-RestartChan:
		if ok == false {
			return nil // 已经关闭
		}
	default:
		// pass
	}

	var args []string

	if len(os.Args) > 1 {
		args = sliceutils.CopySlice(os.Args[1:])
		if !sliceutils.SliceHasItem(args, restartFlag) {
			args = append([]string{restartFlag}, args...)
		}
	} else {
		args = []string{restartFlag}
	}

	cmd := exec.Command(osutils.GetArgs0(), args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: windows.CREATE_NEW_PROCESS_GROUP | windows.CREATE_UNICODE_ENVIRONMENT | windows.CREATE_NEW_CONSOLE,
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	logger.Warnf("the program restart...")

	if err := consoleutils.FreeConsole(); err != nil {
		return err
	}

	logger.Warnf("restart ready")

	close(RestartChan)
	return nil
}

func FromRestart() error {
	logger.Infof("Wait ready...")
	<-time.After(RestartWaitTime)
	logger.Infof("Restart ready...")
	return nil
}

func FirstRun() error {
	err := consoleutils.MakeNewConsole(consoleutils.CodePageUTF8)
	if err != nil {
		return err
	}
	return nil
}
