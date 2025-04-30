// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logpanic"
	"time"
)

func IsReady() bool {
	return internal.IsReady()
}

func Tagf(format string, args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.TagSkipf(1, format, args...)
}

func Debugf(format string, args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	if !internal.IsReady() {
		logpanic.Panic(time.Now(), fmt.Sprintf(format, args...))
	}
	internal.GlobalLogger.Panicf(format, args...)
}

func Tag(args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.TagSkip(1, args...)
}

func Debug(args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Debug(args...)
}

func Info(args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Info(args...)
}

func Warn(args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Warn(args...)
}

func Error(args ...interface{}) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Error(args...)
}

func Panic(args ...interface{}) {
	if !internal.IsReady() {
		logpanic.Panic(time.Now(), fmt.Sprint(args...))
	}
	internal.GlobalLogger.Panic(args...)
}

func TagWrite(msg string) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.TagSkip(1, msg)
}

func DebugWrite(msg string) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.DebugWrite(msg)
}

func InfoWrite(msg string) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.InfoWrite(msg)
}

func WarnWrite(msg string) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.WarnWrite(msg)
}

func ErrorWrite(msg string) {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.ErrorWrite(msg)
}

func PanicWrite(msg string) {
	if !internal.IsReady() {
		logpanic.Panic(time.Now(), msg)
	}
	internal.GlobalLogger.PanicWrite(msg)
}

func Recover() {
	if !internal.IsReady() {
		return
	}
	internal.GlobalLogger.Recover()
}
