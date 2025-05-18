// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitreturn

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
)

const exitCodeSuccess = 0
const exitCodeFailed = 1

var exitcode = exitCodeSuccess

func ReturnError(err error) int {
	templog.TempLogf("error: %s\n", err.Error())
	return exitCodeFailed
}

func ReturnSuccess() int {
	templog.TempLog("success!")
	return exitCodeSuccess
}

func SaveExitCode(_exitcode int) {
	if _exitcode == exitCodeSuccess {
		return // 不做修改
	}
	exitcode = _exitcode
}

func GetExitCode() int {
	return exitcode
}
