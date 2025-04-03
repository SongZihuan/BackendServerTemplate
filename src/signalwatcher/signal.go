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
	var signalList = make([]os.Signal, 0, 4)

	if config.Data().Signal.SigIntExit.IsEnable(true) {
		signalList = append(signalList, syscall.SIGINT)
	}

	if config.Data().Signal.SigTermExit.IsEnable(true) {
		signalList = append(signalList, syscall.SIGTERM)
	}

	if config.Data().Signal.SigHupExit.IsEnable(true) {
		signalList = append(signalList, syscall.SIGHUP)
	}

	if config.Data().Signal.SigQuitExit.IsEnable(false) {
		signalList = append(signalList, syscall.SIGQUIT)
	}

	signal.Notify(exitChannel, signalList...)
	return exitChannel
}
