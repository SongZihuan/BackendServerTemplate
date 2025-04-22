// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !windows

package restart

import (
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/sliceutils"
	"os"
	"os/exec"
	"time"
)

var RestartChan = make(chan bool)

const RestartReadyTime = 5 * time.Second
const RestartExitTime = 5 * time.Second
const RestartWaitTime = RestartReadyTime + RestartExitTime + (3 * time.Second)

func RestartProgram(restartFlag string, beforeReturnHook func()) error {
	select {
	case _, ok := <-RestartChan:
		if ok == false {
			return nil
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
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 保留默认设置，父进程结束后子进程将有init接管
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Setpgid: true, // 确保新的进程组独立
	//}

	if err := cmd.Start(); err != nil {
		return err
	}

	if beforeReturnHook != nil {
		beforeReturnHook()
	}

	close(RestartChan)
	return nil
}

func FromRestart() error {
	<-time.After(RestartWaitTime)
	return nil
}

func FirstRun() error {
	return nil
}
