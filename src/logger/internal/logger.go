// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"sync"
)

var GlobalLogger *Logger = nil

type Logger struct {
	lock       sync.RWMutex
	warnWriter logwriter.Writer
	errWriter  logwriter.Writer
}
