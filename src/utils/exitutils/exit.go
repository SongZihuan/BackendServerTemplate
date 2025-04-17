// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

const (
	exitCodeMin                 = 0
	exitCodeMax                 = 255
	exitCodeErrorLogMustBeReady = 254
)

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

func InitFailedErrorForWin32ConsoleModule(reason string, exitCode ...int) ExitCode {
	if reason == "" {
		reason = "no reason"
	}

	ec := getExitCode(1, exitCode...)

	log.Printf("The module `Win32 Console` init failed (reason: `%s`) .", reason)
	log.Printf("Now we should exit with code %d.", ec)

	return ec
}

func InitFailedErrorForLoggerModule(reason string, exitCode ...int) ExitCode {
	if reason == "" {
		reason = "no reason"
	}

	ec := getExitCode(1, exitCode...)

	log.Printf("The module `Logger` init failed (reason: `%s`) .", reason)
	log.Printf("Now we should exit with code %d.", ec)

	return ec
}

func InitFailedError(module string, reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	if reason == "" {
		reason = "no reason"
	}

	ec := getExitCode(1, exitCode...)

	logger.Errorf("The module `%s` init failed (reason: `%s`) .", module, reason)
	logger.Errorf("Now we should exit with code %d.", ec)

	return ec
}

func RunErrorQuite(exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}
	return getExitCode(1, exitCode...)
}

func RunError(reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	if reason == "" {
		reason = "no reason"
	}

	ec := getExitCode(1, exitCode...)

	logger.Errorf("Run error (reason: `%s`) .", reason)
	logger.Errorf("Now we should exit with code %d.", ec)

	return ec
}

func SuccessExit(reason string, exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	if reason == "" {
		reason = "no reason"
	}

	ec := getExitCode(0, exitCode...)

	logger.Warnf("Now we should exit with code %d (reason: %s) .", ec, reason)

	return ec
}

func SuccessExitSimple(reason string, exitCode ...int) ExitCode {
	if reason != "" {
		fmt.Println(reason)
	}
	return getExitCode(0, exitCode...)
}

func SuccessExitQuite(exitCode ...int) ExitCode {
	if !logger.IsReady() {
		return exitCodeErrorLogMustBeReady
	}

	ec := getExitCode(0, exitCode...)

	return ec
}
