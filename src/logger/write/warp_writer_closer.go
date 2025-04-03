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

func (w *wrapperWriterClose) Close() error {
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
