// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"testing"
)

func TestWrapWriteCloser(t *testing.T) {
	var a logwriter.Writer
	var b *WarpWriteCloser

	a = b
	_ = a
}
