// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"io"
)

func InitBaseLogger(level loglevel.LoggerLevel, logTag bool, humanWarnWriter, humanErrWriter, machineWarnWriter, machineErWriter io.Writer) error {
	return internal.InitLogger(level, logTag, humanWarnWriter, humanErrWriter, machineWarnWriter, machineErWriter)
}

func CloseLogger() {
	internal.CloseLogger()
}
