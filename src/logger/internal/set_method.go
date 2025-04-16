// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/nonewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/wrapwriter"
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

func (l *Logger) SetHumanWarnWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = wrapwriter.WrapToWriter(os.Stdout)
	}

	last := l.humanWarnWriter
	l.humanWarnWriter = w
	return last, nil
}

func (l *Logger) SetHumanErrWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = wrapwriter.WrapToWriter(os.Stderr)
	}

	last := l.humanErrWriter
	l.humanErrWriter = w
	return last, nil
}

func (l *Logger) SetMachineWarnWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = nonewriter.NewNoneWriter()
	}

	last := l.machineWarnWriter
	l.machineWarnWriter = w
	return last, nil
}

func (l *Logger) SetMachineErrWriter(w write.Writer) (write.Writer, error) {
	if w == nil {
		w = nonewriter.NewNoneWriter()
	}

	last := l.machineErrWriter
	l.machineErrWriter = w
	return last, nil
}

func (l *Logger) CloseHumanWarnWriter() error {
	if l.humanWarnWriter == nil {
		return fmt.Errorf("warn writer not set")
	}

	w, ok := l.humanWarnWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.ExitClose()
}

func (l *Logger) CloseHumanErrWriter() error {
	if l.humanErrWriter == nil {
		return fmt.Errorf("error writer not set")
	}

	w, ok := l.humanErrWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.ExitClose()
}

func (l *Logger) CloseMachineWarnWriter() error {
	if l.machineWarnWriter == nil {
		return fmt.Errorf("warn writer not set")
	}

	w, ok := l.machineWarnWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.ExitClose()
}

func (l *Logger) CloseMachineErrWriter() error {
	if l.machineErrWriter == nil {
		return fmt.Errorf("error writer not set")
	}

	w, ok := l.machineErrWriter.(write.WriteCloser)
	if !ok {
		return nil
	}

	return w.ExitClose()
}
