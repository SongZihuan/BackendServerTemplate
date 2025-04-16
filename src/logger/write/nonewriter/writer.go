// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nonewriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
)

type NoneWriter struct {
}

func (f *NoneWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (f *NoneWriter) Close() error {
	return f.ExitClose()
}

func (f *NoneWriter) ExitClose() error {
	return nil
}

func NewNoneWriter() *NoneWriter {
	return new(NoneWriter)
}

func _testNoneWriter() {
	var a write.WriteCloser
	var b *NoneWriter

	a = b
	_ = a
}
