// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wrapwritercloser

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"io"
)

type WrapperWriterClose struct {
	writer io.WriteCloser
}

func (w *WrapperWriterClose) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *WrapperWriterClose) Close() (err error) {
	if w.writer != nil {
		return nil
	}

	defer func() {
		if err == nil {
			w.writer = nil
		}
	}()

	return w.writer.Close()
}

func (w *WrapperWriterClose) ExitClose() error {
	return w.Close()
}

func WraToWriteCloser(w io.WriteCloser) *WrapperWriterClose {
	return &WrapperWriterClose{
		writer: w,
	}
}

func _testWrapWriteCloser() {
	var a write.WriteCloser
	var b *WrapperWriterClose

	a = b
	_ = a
}
