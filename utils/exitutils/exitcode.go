// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"fmt"
	"os"
)

const (
	// 最大值和最小值
	exitCodeMin = 0
	exitCodeMax = 255

	// 默认值
	exitCodeDefaultSuccess = 0 // 默认值：正常
	exitCodeDefaultError   = 1 // 默认值：错误

	// 特殊错误值
	exitCodeInitFailedError  = 2 // 初始化错误
	exitCodeRunError         = 3 // 运行时错误
	exitCodeRunErrorQuite    = 4 // 运行时错误（安静关闭）
	exitCodeStartUpError     = 5
	exitCodeReload           = 252 // 重启信号
	exitCodeWithUnknownError = 253 // 未知错误
	// 254: 原定为 `Logger` 错误的退出码，现已取消。
)

const ExitCodeReload = exitCodeReload

type ExitCode int

func (e ExitCode) Error() string {
	return fmt.Sprintf("Exit with code %d", e)
}

func (e ExitCode) ClampAttribute() ExitCode {
	res := e

	if res < exitCodeMin {
		res = -res
	}

	if res > exitCodeMax {
		res = exitCodeMax
	}

	return res
}

func (e ExitCode) Exit() {
	os.Exit(int(e))
}

func getExitCode(defaultExitCode int, exitCode ...int) (ec ExitCode) {
	if len(exitCode) == 1 {
		ec = ExitCode(exitCode[0])
	} else {
		ec = ExitCode(defaultExitCode)
	}

	return ec.ClampAttribute()
}
