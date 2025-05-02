// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"io"
	"sync"
)

type WarpWriter struct {
	writer io.Writer
	fn     logformat.FormatFunc
	lock   sync.Mutex
}

func (w *WarpWriter) Write(data *logformat.LogData) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.writer == nil {
		return 0, fmt.Errorf("writer is nil")
	}

	return fmt.Fprintf(w.writer, "%s\n", w.fn(data))
}

func NewWarpWriter(w io.Writer, fn logformat.FormatFunc) *WarpWriter {
	if fn == nil {
		fn = logformat.FormatConsole
	}

	return &WarpWriter{
		writer: w,
		fn:     fn,
	}
}
