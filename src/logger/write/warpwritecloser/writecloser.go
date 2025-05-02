// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwritecloser

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/utils/osutils"
	"io"
	"sync"
)

type WarpWriteCloser struct {
	writer io.WriteCloser
	fn     logformat.FormatFunc
	lock   sync.Mutex
}

func (w *WarpWriteCloser) Write(data *logformat.LogData) {
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

func (w *WarpWriteCloser) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	defer func() {
		w.writer = nil
	}()

	return w.writer.Close()
}

func NewWarpWriteCloser(w io.WriteCloser, fn logformat.FormatFunc) *WarpWriteCloser {
	return &WarpWriteCloser{
		writer: w,
		fn:     fn,
	}
}
