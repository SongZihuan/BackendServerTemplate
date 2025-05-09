// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/warpwriter"
	"os"
)

func InitLogger(level loglevel.LoggerLevel, logTag bool, warnWriter, errWriter logwriter.Writer) error {
	logLevel, ok := levelMap[level]
	if !ok {
		return fmt.Errorf("invalid log level: %s", level)
	}

	if warnWriter == nil {
		warnWriter = warpwriter.NewWarpWriter(os.Stdout, nil)
	}

	if errWriter == nil {
		errWriter = warpwriter.NewWarpWriter(os.Stderr, nil)
	}

	logger := &Logger{
		level:      level,
		logLevel:   logLevel,
		logTag:     logTag,
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
