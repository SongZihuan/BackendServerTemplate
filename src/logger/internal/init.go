// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/nonewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/wrapwriter"
	"io"
	"os"
)

func InitLogger(level loglevel.LoggerLevel, logTag bool, humanWarnWriter, humanErrWriter io.Writer, machineWarnWriter, machineErrWriter io.Writer) error {
	logLevel, ok := levelMap[level]
	if !ok {
		return fmt.Errorf("invalid log level: %s", level)
	}

	if humanWarnWriter == nil {
		humanWarnWriter = wrapwriter.WrapToWriter(os.Stdout)
	}

	if humanErrWriter == nil {
		humanErrWriter = wrapwriter.WrapToWriter(os.Stderr)
	}

	if machineWarnWriter == nil {
		machineWarnWriter = nonewriter.NewNoneWriter()
	}

	if machineErrWriter == nil {
		machineErrWriter = nonewriter.NewNoneWriter()
	}

	logger := &Logger{
		level:             level,
		logLevel:          logLevel,
		logTag:            logTag,
		humanWarnWriter:   humanWarnWriter,
		humanErrWriter:    humanErrWriter,
		machineWarnWriter: machineWarnWriter,
		machineErrWriter:  machineErrWriter,
	}

	GlobalLogger = logger
	return nil
}

func IsReady() bool {
	return GlobalLogger != nil
}

func CloseLogger() {
	_ = GlobalLogger.CloseHumanWarnWriter()
	_ = GlobalLogger.CloseHumanErrWriter()
	_ = GlobalLogger.CloseMachineWarnWriter()
	_ = GlobalLogger.CloseMachineErrWriter()
}
