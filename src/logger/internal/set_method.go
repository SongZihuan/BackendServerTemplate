// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/warpwriter"
	"os"
)

func (l *Logger) SetLevel(level loglevel.LoggerLevel) error {
	logLevel, ok := levelMap[level]
	if !ok {
		return fmt.Errorf("invalid log level: %s", level)
	}

	l.level = level
	l.logLevel = logLevel

	return nil
}

func (l *Logger) SetLogTag(logTag bool) error {
	l.logTag = logTag
	return nil
}

func (l *Logger) SetWarnWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = warpwriter.NewWarpWriter(os.Stdout, nil)
	}

	last := l.warnWriter
	l.warnWriter = w
	return last, nil
}

func (l *Logger) SetErrWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = warpwriter.NewWarpWriter(os.Stderr, nil)
	}

	last := l.errWriter
	l.errWriter = w
	return last, nil
}

func (l *Logger) CloseWarnWriter() error {
	if l.warnWriter == nil {
		return fmt.Errorf("warn writer not set")
	}

	w, ok := l.warnWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.Close()
}

func (l *Logger) CloseErrWriter() error {
	if l.errWriter == nil {
		return fmt.Errorf("error writer not set")
	}

	w, ok := l.errWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.Close()
}
