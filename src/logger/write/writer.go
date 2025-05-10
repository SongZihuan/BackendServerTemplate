// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package write

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"io"
)

type Writer interface {
	Write(data *logformat.LogData)
}

type WriteCloser interface {
	Writer
	io.Closer
}
