package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
)

var GlobalLogger *Logger = nil

type Logger struct {
	level             loglevel.LoggerLevel
	logLevel          loggerLevel
	logTag            bool
	humanWarnWriter   write.Writer
	humanErrWriter    write.Writer
	machineWarnWriter write.Writer
	machineErrWriter  write.Writer
	args0Name         string
}
