// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/nonewriter"
)

func (l *Logger) SetWarnWriter(w logwriter.Writer) (logwriter.Writer, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if w == nil {
		w = nonewriter.NewNoneWriter()
	}

	last := l.warnWriter
	l.warnWriter = w
	return last, nil
}

func (l *Logger) SetErrWriter(w logwriter.Writer) (logwriter.Writer, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if w == nil {
		w = nonewriter.NewNoneWriter()
	}

	last := l.errWriter
	l.errWriter = w
	return last, nil
}

func (l *Logger) CloseWriter() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	err1 := l.closeWarnWriter()
	err2 := l.closeErrWriter()

	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	}

	return nil
}

func (l *Logger) CloseWarnWriter() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.closeWarnWriter()
}

func (l *Logger) CloseErrWriter() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.closeErrWriter()
}

func (l *Logger) closeWarnWriter() error {
	if l.warnWriter == nil {
		return fmt.Errorf("warn writer not set")
	}

	w, ok := l.warnWriter.(logwriter.Writer)
	if !ok {
		return nil
	}

	return w.Close()
}

func (l *Logger) closeErrWriter() error {
	if l.errWriter == nil {
		return fmt.Errorf("error writer not set")
	}

	w, ok := l.errWriter.(logwriter.Writer)
	if !ok {
		return nil
	}

	return w.Close()
}
