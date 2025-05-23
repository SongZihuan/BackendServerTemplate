// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logpanic"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"time"
)

func (l *Logger) TagSkipf(skip int, format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	content := fmt.Sprintf(format, args...)
	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(rtdata.GetLocation())

	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprintf(format, args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) TagSkip(skip int, args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	content := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())
	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) Debug(args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) Info(args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) Warn(args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) Error(args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) Panic(args ...interface{}) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	msg := fmt.Sprint(args...)
	now := time.Now().In(rtdata.GetLocation())

	// 此处必须使用 <- 确保输出完成，然后才能 panic
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) TagSkipWrite(skip int, content string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.PseudoLevelTag, msg, now))
}

func (l *Logger) DebugWrite(msg string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelDebug, msg, now))
}

func (l *Logger) InfoWrite(msg string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelInfo, msg, now))
}

func (l *Logger) WarnWrite(msg string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	now := time.Now().In(rtdata.GetLocation())
	<-l.warnWriter.Write(logformat.GetLogData(loglevel.LevelWarn, msg, now))
}

func (l *Logger) ErrorWrite(msg string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	now := time.Now().In(rtdata.GetLocation())
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelError, msg, now))
}

func (l *Logger) PanicWrite(msg string) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	now := time.Now().In(rtdata.GetLocation())
	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))

	logpanic.Panic(now, msg)
}

func (l *Logger) Recover() {
	err := recover()
	if err == nil {
		return
	}

	msg := ""
	now := time.Now().In(rtdata.GetLocation())

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

	l.lock.RLock()
	defer l.lock.RUnlock()

	<-l.errWriter.Write(logformat.GetLogData(loglevel.LevelPanic, msg, now))
}
