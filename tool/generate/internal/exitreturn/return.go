// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitreturn

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
)

const exitCodeSuccess = 0
const exitCodeFailed = 1

func ReturnError(err error) int {
	genlog.GenLogf("error: %s\n", err.Error())
	return exitCodeFailed
}

func ReturnSuccess() int {
	genlog.GenLog("success!")
	return exitCodeSuccess
}
