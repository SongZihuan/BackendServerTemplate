// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package write

import "io"

type Writer interface {
	io.Writer
}

type WriteCloser interface {
	io.WriteCloser
	ExitClose() error
}
