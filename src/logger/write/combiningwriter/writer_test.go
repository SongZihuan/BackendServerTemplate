// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package combiningwriter

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"testing"
)

func TestCombiningWriter(t *testing.T) {
	var a write.WriteCloser
	var b *CombiningWriter

	a = b
	_ = a
}
