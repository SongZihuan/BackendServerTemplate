// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

func SuccessExit(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeDefaultSuccess, exitCode...)

	if logger.IsReady() {
		if reason != "" {
			logger.Warnf("Exit %d: %s", ec, reason)
		} else {
			logger.Warnf("Exit %d: ok", ec)
		}
	} else {
		if reason != "" {
			log.Printf("Exit %d: %s\n", ec, reason)
		} else {
			log.Printf("Exit %d: ok\n", ec)
		}
	}

	return ec // ec 不再受 logger 的 ready 问题影响
}
