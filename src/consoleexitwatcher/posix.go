// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !windows

package consoleexitwatcher

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/utils/consoleutils"
)

var consoleexitchan = make(chan any)
var consolewaitexitchan = make(chan any)

func NewWin32ConsoleExitChannel() (chan any, chan any, error) {
	return consoleexitchan, consolewaitexitchan, nil
}

func GetExitEvent() consoleutils.Event {
	logger.Panicf("not event")
	return nil
}
