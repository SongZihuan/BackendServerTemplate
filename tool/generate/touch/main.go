// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/touch"
	"os"
)

// 冗余导入此包，该包包含必须导入的全部信息
import (
	_ "github.com/SongZihuan/BackendServerTemplate/global/pkgimport"
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
