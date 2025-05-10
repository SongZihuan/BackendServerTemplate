// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/warpwriter"
	"os"
)

func InitLogger(level loglevel.LoggerLevel, tag bool, warnWriter, errWriter logwriter.Writer) error {
	if warnWriter == nil {
		warnWriter = warpwriter.NewWarpWriter(level, tag, os.Stdout, logformat.FormatConsole)
	}

	if errWriter == nil {
		errWriter = warpwriter.NewWarpWriter(level, tag, os.Stderr, logformat.FormatConsole)
	}

	logger := &Logger{
		warnWriter: warnWriter,
		errWriter:  errWriter,
	}

	GlobalLogger = logger
	return nil
}

func IsReady() bool {
	return GlobalLogger != nil
}

func CloseLogger() {
	_ = GlobalLogger.CloseWriter()
	GlobalLogger = nil
}
