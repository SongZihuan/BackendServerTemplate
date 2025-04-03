// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/runtimeutils"
)

func (l *Logger) Executablef(format string, args ...interface{}) string {
	str := fmt.Sprintf(format, args...)
	if str == "" {
		_, _ = fmt.Fprintf(l.warnWriter, "[Executable]: %s\n", l.args0)
	} else {
		_, _ = fmt.Fprintf(l.warnWriter, "[Executable %s]: %s\n", l.args0, str)
	}
	return l.args0
}

func (l *Logger) Tagf(format string, args ...interface{}) {
	l.TagSkipf(1, format, args...)
}

func (l *Logger) TagSkipf(skip int, format string, args ...interface{}) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Tag %s]: %s %s %s:%d\n", l.args0Name, str, funcName, file, line)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Debug %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Info %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Warning %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.errWriter, "[Error %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	str := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.errWriter, "[Panic %s]: %s\n", l.args0Name, str)

	if l.realPanic {
		panic(str)
	}
}

func (l *Logger) Tag(args ...interface{}) {
	l.TagSkip(1, args...)
}

func (l *Logger) TagSkip(skip int, args ...interface{}) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Tag %s]: %s %s %s:%d\n", l.args0Name, str, funcName, file, line)
}

func (l *Logger) Debug(args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Debug %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Info(args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Info %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.warnWriter, "[Warning %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Error(args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.errWriter, "[Error %s]: %s\n", l.args0Name, str)
}

func (l *Logger) Panic(args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	str := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.errWriter, "[Panic %s]: %s\n", l.args0Name, str)

	if l.realPanic {
		panic(str)
	}
}

func (l *Logger) TagWrite(msg string) {
	l.TagSkipWrite(1, msg)
}

func (l *Logger) TagSkipWrite(skip int, msg string) {
	if !l.logTag {
		return
	}

	funcName, file, _, line := runtimeutils.GetCallingFunctionInfo(skip + 1)

	_, _ = fmt.Fprintf(l.warnWriter, "[Debug %s]: %s %s %s:%d\n", l.args0Name, msg, funcName, file, line)
}

func (l *Logger) DebugWrite(msg string) {
	if l.logLevel >= levelDebug {
		return
	}

	_, _ = fmt.Fprintf(l.warnWriter, "[Debug %s]: %s\n", l.args0Name, msg)
}

func (l *Logger) InfoWrite(msg string) {
	if l.logLevel >= levelInfo {
		return
	}

	_, _ = fmt.Fprintf(l.warnWriter, "[Info %s]: %s\n", l.args0Name, msg)
}

func (l *Logger) WarnWrite(msg string) {
	if l.logLevel >= levelWarn {
		return
	}

	_, _ = fmt.Fprintf(l.warnWriter, "[Warning %s]: %s\n", l.args0Name, msg)
}

func (l *Logger) ErrorWrite(msg string) {
	if l.logLevel >= levelError {
		return
	}

	_, _ = fmt.Fprintf(l.errWriter, "[Error %s]: %s\n", l.args0Name, msg)
}

func (l *Logger) PanicWrite(msg string) {
	if l.logLevel >= levelPanic {
		return
	}

	_, _ = fmt.Fprintf(l.errWriter, "[Panic %s]: %s\n", l.args0Name, msg)

	if l.realPanic {
		panic(msg)
	}
}
