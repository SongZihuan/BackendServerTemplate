// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/utils/osutils"
	"io"
	"sync"
)

type WarpWriter struct {
	writer io.Writer
	fn     logformat.FormatFunc
	lock   sync.Mutex
}

func (w *WarpWriter) Write(data *logformat.LogData) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.writer == nil {
		return
	}

	_, _ = fmt.Fprintf(w.writer, "%s\n", w.fn(data))

	syncer, ok := w.writer.(osutils.Syncer)
	if ok {
		_ = syncer.Sync()
	}
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
