// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"io"
	"os"
)

func InitLogger(level loglevel.LoggerLevel, logTag bool, realPanic bool, warnWriter, errWriter io.Writer) error {
	logLevel, ok := levelMap[level]
	if !ok {
		return fmt.Errorf("invalid log level: %s", level)
	}

	if warnWriter == nil {
		warnWriter = write.ChangeToWriter(os.Stdout)
	}

	if errWriter == nil {
		errWriter = write.ChangeToWriter(os.Stderr)
	}

	logger := &Logger{
		level:      level,
		logLevel:   logLevel,
		logTag:     logTag,
		warnWriter: warnWriter,
		errWriter:  errWriter,
		args0:      osutils.GetArgs0(),
		args0Name:  osutils.GetArgs0Name(),
		realPanic:  realPanic,
	}

	GlobalLogger = logger
	return nil
}

func IsReady() bool {
	return GlobalLogger != nil
}

func CloseLogger() {

}
