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

type WarpWrite struct {
	writer io.Writer
	close  bool // 表示 warp 关闭时，是否关闭 writer
	fn     logformat.FormatFunc
	mutex  sync.Mutex
}

func (w *WarpWrite) Write(data *logformat.LogData) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.writer == nil {
		return
	}

	_, _ = fmt.Fprintf(w.writer, "%s\n", w.fn(data))

	syncer, ok := w.writer.(osutils.Syncer)
	if ok {
		_ = syncer.Sync()
	}
}

func (w *WarpWrite) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	defer func() {
		w.writer = nil
		w.close = false
	}()

	if closer, ok := w.writer.(io.WriteCloser); w.close && ok {
		return closer.Close()
	}

	return nil
}

func NewWarpWriter(w io.Writer, fn logformat.FormatFunc) *WarpWrite {
	return &WarpWrite{
		writer: w,
		fn:     fn,
		close:  false,
	}
}

func NewWarpWriteCloser(w io.Writer, fn logformat.FormatFunc) *WarpWrite {
	return &WarpWrite{
		writer: w,
		fn:     fn,
		close:  true,
	}
}
