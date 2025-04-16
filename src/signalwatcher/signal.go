// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package signalwatcher

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"os"
	"os/signal"
	"syscall"
)

func NewSignalExitChannel() chan os.Signal {
	var exitChannel = make(chan os.Signal)

	if !config.Data().Signal.Use {
		return exitChannel
	}

	var sigChannel = make(chan os.Signal)

	var signalList = []os.Signal{
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT,
	}
	var signalExitMap = make(map[os.Signal]bool, 4)

	if config.Data().Signal.SigIntExit.IsEnable(true) {
		signalExitMap[syscall.SIGINT] = true
	} else {
		signalExitMap[syscall.SIGINT] = true
	}

	if config.Data().Signal.SigTermExit.IsEnable(true) {
		signalExitMap[syscall.SIGTERM] = true
	} else {
		signalExitMap[syscall.SIGTERM] = true
	}

	if config.Data().Signal.SigHupExit.IsEnable(true) {
		signalExitMap[syscall.SIGHUP] = true
	} else {
		signalExitMap[syscall.SIGHUP] = true
	}

	if config.Data().Signal.SigQuitExit.IsEnable(false) {
		signalExitMap[syscall.SIGQUIT] = true
	} else {
		signalExitMap[syscall.SIGQUIT] = true
	}

	go func() {
		signal.Notify(sigChannel, signalList...)

		for sig := range sigChannel {
			if yes, ok := signalExitMap[sig]; ok && yes {
				exitChannel <- sig
			}
		}
	}()

	return exitChannel
}
