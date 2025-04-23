// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
)

func InitBaseLogger() error {
	if IsReady() {
		return nil
	}
	return internal.InitLogger(loglevel.LevelDebug, true, nil, nil)
}

func CloseLogger() {
	if !IsReady() {
		return
	}
	internal.CloseLogger()
}
