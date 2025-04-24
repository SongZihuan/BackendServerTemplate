// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package modutils

import (
	"runtime/debug"
	"strings"
)

var ModPath = ""
var IsGitHub bool = false

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("read build info failed")
	}

	ModPath = info.Main.Path
	IsGitHub = strings.HasPrefix(ModPath, "github.com/")
}
