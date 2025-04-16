// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/mattn/go-isatty"
	"os"
)

func (l *Logger) GetLevel() loglevel.LoggerLevel {
	return l.level
}

func (l *Logger) IsLogTag() bool {
	return l.logTag
}

func (l *Logger) IsWarnWriterTerm() bool {
	w, ok := l.humanWarnWriter.(*os.File)
	if !ok {
		return false
	} else if !isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()) { // 非终端
		return false
	}
	return true
}

func (l *Logger) IsErrWriterTerm() bool {
	w, ok := l.humanErrWriter.(*os.File)
	if !ok {
		return false
	} else if !isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()) { // 非终端
		return false
	}
	return true
}

func (l *Logger) IsTermDump() bool {
	// TERM为dump表示终端为基础模式，不支持高级显示
	return os.Getenv("TERM") == "dumb"
}

func (l *Logger) IsWarnWriterTermNoDump() bool {
	return l.IsWarnWriterTerm() && !l.IsTermDump()
}

func (l *Logger) IsErrWriterTermNoDump() bool {
	return l.IsErrWriterTerm() && !l.IsTermDump()
}
