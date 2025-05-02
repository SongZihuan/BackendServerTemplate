// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwritecloser

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"io"
	"sync"
)

type WarpWriteCloser struct {
	writer io.WriteCloser
	fn     logformat.FormatFunc
	lock   sync.Mutex
}

func (w *WarpWriteCloser) Write(data *logformat.LogData) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.writer == nil {
		return 0, fmt.Errorf("writer has been close")
	}

	return fmt.Fprintf(w.writer, "%s\n", w.fn(data))
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
