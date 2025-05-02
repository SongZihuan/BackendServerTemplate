// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"sync"
)

var GlobalLogger *Logger = nil

type Logger struct {
	lock       sync.RWMutex
	level      loglevel.LoggerLevel
	logLevel   loggerLevel
	logTag     bool
	warnWriter write.Writer
	errWriter  write.Writer
	args0Name  string
}
