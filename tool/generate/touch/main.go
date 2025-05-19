// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/touch"
	"os"
)

func main() {
	os.Exit(command())
}

// command 单独把 touch 作为一个程序，避免程序本身对 resource 包的依赖。
func command() (exitcode int) {
	err := touch.TouchBaseFile()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
