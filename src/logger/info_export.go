// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
)

func GetLevel() loglevel.LoggerLevel {
	if !internal.IsReady() {
		return loglevel.LevelDebug
	}
	return internal.GlobalLogger.GetLevel()
}

func IsLogTag() bool {
	if !internal.IsReady() {
		return false
	}
	return internal.GlobalLogger.IsLogTag()
}
