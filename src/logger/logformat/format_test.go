// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logformat

import "testing"

func TestFormat(t *testing.T) {
	var a FormatFunc

	a = FormatMachine
	a = FormatFile
	a = FormatConsole
	a = FormatConsolePretty

	_ = a
}
