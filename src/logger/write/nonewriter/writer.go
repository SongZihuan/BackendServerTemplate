// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nonewriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
)

type NoneWriter struct {
}

func (f *NoneWriter) Write(_ *logformat.LogData) {}

func (f *NoneWriter) Close() error {
	return nil
}

func NewNoneWriter() *NoneWriter {
	return new(NoneWriter)
}
