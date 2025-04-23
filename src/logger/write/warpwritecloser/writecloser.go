// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwritecloser

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"io"
)

type WarpWriteCloser struct {
	writer io.WriteCloser
	fn     logformat.FormatFunc
}

func (w *WarpWriteCloser) Write(data *logformat.LogData) (n int, err error) {
	return fmt.Fprintf(w.writer, "%s\n", w.fn(data))
}

func (w *WarpWriteCloser) Close() error {
	return w.writer.Close()
}

func NewWarpWriteCloser(w io.WriteCloser, fn logformat.FormatFunc) *WarpWriteCloser {
	return &WarpWriteCloser{
		writer: w,
		fn:     fn,
	}
}

func _testWrapWriteCloser() {
	var a write.Writer
	var b *WarpWriteCloser

	a = b
	_ = a
}
