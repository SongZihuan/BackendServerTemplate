// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
)

func (l *Logger) GetLevel() loglevel.LoggerLevel {
	return l.level
}

func (l *Logger) IsLogTag() bool {
	return l.logTag
}
