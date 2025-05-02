// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package osutils

import (
	"os"
	"testing"
)

func TestStdout(t *testing.T) {
	var a Syncer
	a = os.Stdout
	_ = a
}

func TestStderr(t *testing.T) {
	var a Syncer
	a = os.Stderr
	_ = a
}
