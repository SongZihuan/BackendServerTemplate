// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sigexitwatcher

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	SIGINT  = syscall.SIGINT
	SIGTERM = syscall.SIGTERM
	SIGHUP  = syscall.SIGHUP
	SIGQUIT = syscall.SIGQUIT
)

var notifySignalList = []os.Signal{SIGINT, SIGTERM, SIGHUP, SIGQUIT}
var sigexitchan = make(chan any)
var lastSignal os.Signal = nil
var once sync.Once

func GetSignalExitChannelFromConfig() chan any {
	return GetSignalExitChannel(getSignalOfConcernFromConfig())
}

func GetSignalExitChannel(concernMap map[os.Signal]bool) chan any {
	once.Do(func() {
		if !config.Data().Signal.Use {
			return
		}

		go func() {
			notifyChannel := make(chan os.Signal, 5)
			signal.Notify(notifyChannel, notifySignalList...)

		SignalNotifyCycle:
			for sig := range notifyChannel {
				if yes, ok := concernMap[sig]; ok && yes {
					lastSignal = sig
					signal.Stop(notifyChannel)
					close(notifyChannel)
					close(sigexitchan)
					break SignalNotifyCycle
				}
			}
		}()
	})

	return sigexitchan
}

func GetExitSignal() os.Signal {
	if lastSignal == nil {
		logger.Panicf("not signal")
	}

	return lastSignal
}

func getSignalOfConcernFromConfig() map[os.Signal]bool {
	var concernMap = make(map[os.Signal]bool, 4)

	if config.Data().Signal.SigIntExit.IsEnable(true) {
		concernMap[SIGINT] = true
	} else {
		concernMap[SIGINT] = false
	}

	if config.Data().Signal.SigTermExit.IsEnable(true) {
		concernMap[SIGTERM] = true
	} else {
		concernMap[SIGTERM] = false
	}

	if config.Data().Signal.SigHupExit.IsEnable(true) {
		concernMap[SIGHUP] = true
	} else {
		concernMap[SIGHUP] = false
	}

	if config.Data().Signal.SigQuitExit.IsEnable(false) {
		concernMap[SIGQUIT] = true
	} else {
		concernMap[SIGQUIT] = false
	}

	return concernMap
}
