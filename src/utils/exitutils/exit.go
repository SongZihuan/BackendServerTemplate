// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
	"os"
)

const (
	exitCodeMin                 = 0
	exitCodeMax                 = 255
	exitCodeDefaultSuccess      = 0   // 默认值：正常
	exitCodeDefaultError        = 1   // 默认值：错误
	exitCodeInitFailedError     = 2   // 初始化错误
	exitCodeRunError            = 3   // 运行时错误
	exitCodeRunErrorQuite       = 4   // 运行时错误（安静关闭）
	exitCodeReload              = 252 // 重启信号
	exitCodeWithUnknownError    = 253 // 未知错误
	exitCodeErrorLogMustBeReady = 254 // 报告该错误需要日志系统加载完成
)

const ExitCodeReload = exitCodeReload

type ExitCode int

func (e ExitCode) Error() string {
	return fmt.Sprintf("Exit with code %d", e)
}

func getExitCode(defaultExitCode int, exitCode ...int) (ec ExitCode) {
	if len(exitCode) == 1 {
		ec = ExitCode(exitCode[0])
	} else {
		ec = ExitCode(defaultExitCode)
	}

	if ec < exitCodeMin {
		ec = -ec
	}

	if ec > exitCodeMax {
		ec = exitCodeMax
	}

	return ec
}

func initModuleFailedLog(module string, reason string) string {
	if module == "" {
		panic("module can not be empty")
	}

	if reason != "" {
		return fmt.Sprintf("Init failed [ %s ]: %s", module, reason)
	} else {
		return fmt.Sprintf("Init failed [ %s ]", module)
	}
}

func InitFailedForWin32ConsoleModule(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeInitFailedError, exitCode...)

	log.Printf(initModuleFailedLog("Win32 Console API", reason))
	log.Printf("Init error exit %d: failed", ec)

	return ec
}

func InitFailedForTimeLocationModule(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeInitFailedError, exitCode...)

	log.Printf(initModuleFailedLog("Time Location", reason))
	log.Printf("Init error exit %d: failed", ec)

	return ec
}

func InitFailedForLoggerModule(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeInitFailedError, exitCode...)

	log.Printf(initModuleFailedLog("Logger", reason))
	log.Printf("Init error exit %d: failed", ec)

	return ec
}

func InitFailed(module string, reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	ec := getExitCode(exitCodeInitFailedError, exitCode...)

	logger.Error(initModuleFailedLog(module, reason))
	logger.Errorf("Init error exit %d: failed", ec)

	return ec
}

func RunErrorQuite(exitCode ...int) ExitCode {
	return getExitCode(exitCodeRunErrorQuite, exitCode...)
}

func RunError(reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	ec := getExitCode(exitCodeRunError, exitCode...)

	if reason != "" {
		logger.Errorf("Run error exit %d: %s", ec, reason)
	} else {
		logger.Errorf("Run error exit %d: failed", ec)
	}

	return ec
}

func SuccessExit(reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	ec := getExitCode(exitCodeDefaultSuccess, exitCode...)

	if reason != "" {
		logger.Warnf("Exit %d: %s", ec, reason)
	} else {
		logger.Warnf("Exit %d: ok", ec)
	}

	return ec
}

func Exit(err error) {
	var ec ExitCode
	if err == nil {
		os.Exit(exitCodeDefaultSuccess)
	} else if errors.As(err, &ec) {
		ExitByCode(ec)
	} else {
		if logger.IsReady() {
			logger.Errorf("Exit %d: %s", ec, err.Error())
		} else {
			log.Printf("Exit %d: %s\n", ec, err.Error())
		}
		os.Exit(exitCodeDefaultError)
	}
}

func ExitQuite(err error) {
	var ec ExitCode
	if err == nil {
		os.Exit(exitCodeDefaultSuccess)
	} else if errors.As(err, &ec) {
		ExitByCode(ec)
	}
	os.Exit(exitCodeDefaultError)
}

func ExitByCode(ec ExitCode) {
	os.Exit(int(ec))
}
