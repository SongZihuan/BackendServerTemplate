// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import "github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"

type loggerLevel int64

const (
	levelDebug loggerLevel = 1
	levelInfo  loggerLevel = 2
	levelWarn  loggerLevel = 3
	levelError loggerLevel = 4
	levelPanic loggerLevel = 5
	levelNone  loggerLevel = 6
)

var levelMap = map[loglevel.LoggerLevel]loggerLevel{
	loglevel.LevelDebug: levelDebug,
	loglevel.LevelInfo:  levelInfo,
	loglevel.LevelWarn:  levelWarn,
	loglevel.LevelError: levelError,
	loglevel.LevelPanic: levelPanic,
	loglevel.LevelNone:  levelNone,
}
