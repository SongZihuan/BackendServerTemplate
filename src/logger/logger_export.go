package logger

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
)

func IsReady() bool {
	return internal.IsReady()
}

func Executablef(format string, args ...interface{}) string {
	if !internal.IsReady() {
		return ""
	}
	return internal.GlobalLogger.Executablef(format, args...)
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
		return
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
		return
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
		return
	}
	internal.GlobalLogger.PanicWrite(msg)
}
