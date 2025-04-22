// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package restart

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"os"
	"os/exec"
)

// 以下为 Sub 函数

var RestartChan = make(chan any)

func SetRestart() {
	select {
	case <-RestartChan:
		return
	default:
		close(RestartChan)
	}
}

// 以下为 Parent 函数

func RunRestart() chan any {
	stopchan := make(chan any)
	go Restart(stopchan)
	return stopchan
}

func Restart(stopchan chan any) {
	var args []string
	if len(os.Args) > 1 {
		args = append([]string{fmt.Sprintf("%s=%d", restartFlag, os.Getpid())}, os.Args[1:]...)
	} else {
		args = []string{fmt.Sprintf("%s=%d", restartFlag, os.Getpid())}
	}

RestartCycle:
	for {
		stop := restart(args)
		if stop {
			break RestartCycle
		}
	}

	close(stopchan)
}

func restart(args []string) (stop bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("restart program error: %v", err)
			stop = false
		}
	}()

	cmd := exec.Command(osutils.GetArgs0(), args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		logger.Errorf("restart program start failed: %s", err.Error())
		return false
	}

	_ = cmd.Wait()
	ec := cmd.ProcessState.ExitCode()

	if ec == exitutils.ExitCodeReload {
		return false
	}
	return true
}
