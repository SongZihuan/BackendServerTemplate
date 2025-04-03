// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package write

import "io"

type wrapperWriter struct {
	writer io.Writer
}

func (w *wrapperWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func ChangeToWriter(w io.Writer) Writer {
	return &wrapperWriter{
		writer: w,
	}
}
