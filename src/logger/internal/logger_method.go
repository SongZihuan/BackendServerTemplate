// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/runtimeutils"
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
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman("TAG", msg))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelDebug, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelDebug, msg))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelInfo, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelInfo, msg))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelWarn, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelWarn, msg))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.humanErrWriter, "%s", l.formatHuman(loglevel.LevelError, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelError, msg))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.humanErrWriter, "%s", l.formatHuman(loglevel.LevelPanic, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelPanic, msg))

	panic(msg)
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
	msg := fmt.Sprintf("%s %s %s:%d", content, funcName, file, line)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman("TAG", msg))
}

func (l *Logger) Debug(args ...interface{}) {
	if l.logLevel >= levelDebug {
		return
	}

	msg := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelDebug, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelDebug, msg))
}

func (l *Logger) Info(args ...interface{}) {
	if l.logLevel >= levelInfo {
		return
	}

	msg := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelInfo, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelInfo, msg))
}

func (l *Logger) Warn(args ...interface{}) {
	if l.logLevel >= levelWarn {
		return
	}

	msg := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelWarn, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelWarn, msg))
}

func (l *Logger) Error(args ...interface{}) {
	if l.logLevel >= levelError {
		return
	}

	msg := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.humanErrWriter, "%s", l.formatHuman(loglevel.LevelError, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelError, msg))
}

func (l *Logger) Panic(args ...interface{}) {
	if l.logLevel >= levelPanic {
		return
	}

	msg := fmt.Sprint(args...)
	_, _ = fmt.Fprintf(l.humanErrWriter, "%s", l.formatHuman(loglevel.LevelPanic, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelPanic, msg))

	panic(msg)
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
	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman("TAG", msg))
}

func (l *Logger) DebugWrite(msg string) {
	if l.logLevel >= levelDebug {
		return
	}

	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelDebug, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelDebug, msg))
}

func (l *Logger) InfoWrite(msg string) {
	if l.logLevel >= levelInfo {
		return
	}

	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelInfo, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelInfo, msg))
}

func (l *Logger) WarnWrite(msg string) {
	if l.logLevel >= levelWarn {
		return
	}

	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelWarn, msg))
	_, _ = fmt.Fprintf(l.machineWarnWriter, "%s", l.formatMachine(loglevel.LevelWarn, msg))
}

func (l *Logger) ErrorWrite(msg string) {
	if l.logLevel >= levelError {
		return
	}

	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelError, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelError, msg))
}

func (l *Logger) PanicWrite(msg string) {
	if l.logLevel >= levelPanic {
		return
	}

	_, _ = fmt.Fprintf(l.humanWarnWriter, "%s", l.formatHuman(loglevel.LevelPanic, msg))
	_, _ = fmt.Fprintf(l.machineErrWriter, "%s", l.formatMachine(loglevel.LevelPanic, msg))

	panic(msg)
}
