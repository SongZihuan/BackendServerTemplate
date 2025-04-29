// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

func RunErrorQuite(exitCode ...int) ExitCode {
	return getExitCode(exitCodeRunErrorQuite, exitCode...)
}

func RunError(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeRunError, exitCode...)

	if logger.IsReady() {
		if reason != "" {
			logger.Errorf("Run error exit %d: %s", ec, reason)
		} else {
			logger.Errorf("Run error exit %d: failed", ec)
		}
		return ec
	} else {
		if reason != "" {
			log.Printf("Run error exit %d: %s\n", ec, reason)
		} else {
			log.Printf("Run error exit %d: failed\n", ec)
		}
		return ec
	}
}
