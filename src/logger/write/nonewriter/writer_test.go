// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nonewriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"testing"
)

func TestNoneWriter(t *testing.T) {
	var a write.WriteCloser
	var b *NoneWriter

	a = b
	_ = a
}
