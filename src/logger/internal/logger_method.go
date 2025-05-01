// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logpanic"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/runtimeutils"
	"time"
)

func (l *Logger) Tagf(format string, args ...interface{}) {
	l.TagSkipf(1, format, args...)
}

func (l *Logger) TagSkipf(skip int, format string, args ...interface{}) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	content := fmt.Sprintf(format, args...)
	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(global.Location)
	_, _ = l.errWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(global.Location)
	_, _ = l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) Tag(args ...interface{}) {
	l.TagSkip(1, args...)
}

func (l *Logger) TagSkip(skip int, args ...interface{}) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	content := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) Debug(args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	msg := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) Info(args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	msg := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) Warn(args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	msg := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) Error(args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	msg := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	_, _ = l.errWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) Panic(args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	msg := fmt.Sprint(args...)
	now := time.Now().In(global.Location)
	_, _ = l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) TagWrite(msg string) {
	l.TagSkipWrite(1, msg)
}

func (l *Logger) TagSkipWrite(skip int, content string) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) DebugWrite(msg string) {
	if l.logLevel >= levelDebug {
		return
	}

	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) InfoWrite(msg string) {
	if l.logLevel >= levelInfo {
		return
	}

	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) WarnWrite(msg string) {
	if l.logLevel >= levelWarn {
		return
	}

	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) ErrorWrite(msg string) {
	if l.logLevel >= levelError {
		return
	}

	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) PanicWrite(msg string) {
	if l.logLevel >= levelPanic {
		return
	}

	now := time.Now().In(global.Location)
	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) Recover() {
	err := recover()
	if err == nil {
		return
	}

	msg := ""
	now := time.Now().In(global.Location)

	if _, ok := err.(*logpanic.PanicData); ok { // 如果是: PanicData，则表示已经输出日志，不需要再输出
		return
	}

	if str, ok := err.(fmt.Stringer); ok {
		msg = str.String()
	} else if _err, ok := err.(error); ok {
		msg = _err.Error()
	} else {
		msg = fmt.Sprintf("%v", err)
	}

	_, _ = l.warnWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))
}
