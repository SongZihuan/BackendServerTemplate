package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
)

var GlobalLogger *Logger = nil

type Logger struct {
	level      loglevel.LoggerLevel
	logLevel   loggerLevel
	logTag     bool
	warnWriter write.Writer
	errWriter  write.Writer
	args0      string
	args0Name  string
	realPanic  bool
}
