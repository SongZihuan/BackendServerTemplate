// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

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

func InitFailed(module string, reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeInitFailedError, exitCode...)

	if logger.IsReady() {
		logger.Error(initModuleFailedLog(module, reason))
		logger.Errorf("Init error exit %d: failed", ec)
		return ec
	} else {
		log.Println(initModuleFailedLog(module, reason))
		log.Printf("Init error exit %d: failed\n", ec)
		return ec
	}
}
