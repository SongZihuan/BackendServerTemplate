// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nonewriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
)

type NoneWriter struct {
}

func (f *NoneWriter) Write(_ *logformat.LogData) (n int, err error) {
	return 0, nil
}

func (f *NoneWriter) Close() error {
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
