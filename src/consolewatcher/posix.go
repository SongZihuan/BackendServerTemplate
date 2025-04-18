// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !windows

package consolewatcher

import "github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"

func NewWin32ConsoleExitChannel() (chan consoleutils.Event, chan any, error) {
	var exitChannel = make(chan consoleutils.Event)
	var waitExitChannel = make(chan any)

	return exitChannel, waitExitChannel, nil
}
