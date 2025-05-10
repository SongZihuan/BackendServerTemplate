// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package loglevel

import (
	"math"
	"strings"
)

type LoggerLevel string

const (
	LevelDebug     LoggerLevel = "debug"
	LevelInfo      LoggerLevel = "info"
	LevelWarn      LoggerLevel = "warn"
	LevelError     LoggerLevel = "error"
	LevelPanic     LoggerLevel = "panic"
	LevelNone      LoggerLevel = "none"
	PseudoLevelTag LoggerLevel = "tag"
)

var LoggerLevelMap = map[LoggerLevel]int{
	LevelDebug:     1,
	LevelInfo:      2,
	LevelWarn:      3,
	LevelError:     4,
	LevelPanic:     5,
	LevelNone:      6,
	PseudoLevelTag: math.MaxInt,
}

func (level LoggerLevel) ToLower() LoggerLevel {
	if !level.OK() {
		return level
	}

	return LoggerLevel(strings.ToLower(string(level)))
}

func (level LoggerLevel) Int() int {
	res, ok := LoggerLevelMap[level]
	if !ok {
		panic("unknown logger level")
	}
	return res
}

func (level LoggerLevel) IntQuite() int {
	res, ok := LoggerLevelMap[level]
	if !ok {
		return 0
	}
	return res
}

func (level LoggerLevel) OK() bool {
	_, ok := LoggerLevelMap[level]
	return ok
}
