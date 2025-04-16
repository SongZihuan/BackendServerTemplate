// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wrapwriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"io"
)

type WrapWriter struct {
	writer io.Writer
}

func (w *WrapWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func WrapToWriter(w io.Writer) *WrapWriter {
	return &WrapWriter{
		writer: w,
	}
}

func _testWrapWriter() {
	var a write.Writer
	var b *WrapWriter

	a = b
	_ = a
}
