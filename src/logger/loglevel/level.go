// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package loglevel

type LoggerLevel string

const PseudoLevelTag LoggerLevel = "TAG"

const (
	LevelDebug LoggerLevel = "debug"
	LevelInfo  LoggerLevel = "info"
	LevelWarn  LoggerLevel = "warn"
	LevelError LoggerLevel = "error"
	LevelPanic LoggerLevel = "panic"
	LevelNone  LoggerLevel = "none"
)

var LoggerLevelMap = map[LoggerLevel]bool{
	LevelDebug: true,
	LevelInfo:  true,
	LevelWarn:  true,
	LevelError: true,
	LevelPanic: true,
	LevelNone:  true,
}
