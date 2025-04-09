// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package write

import "io"

type wrapperWriterClose struct {
	writer io.WriteCloser
}

func (w *wrapperWriterClose) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *wrapperWriterClose) Close() (err error) {
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

func (w *wrapperWriterClose) ExitClose() error {
	return w.Close()
}

func ChangeToWriteCloser(w io.WriteCloser) WriteCloser {
	return &wrapperWriterClose{
		writer: w,
	}
}

func _testWarpWriteCloser() {
	var a *wrapperWriterClose
	var b WriteCloser

	b = a
	_ = b
}
