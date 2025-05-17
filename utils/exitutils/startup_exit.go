// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

func StartupError(reason string, exitCode ...int) ExitCode {
	ec := getExitCode(exitCodeStartUpError, exitCode...)

	if logger.IsReady() {
		if reason != "" {
			logger.Errorf("Start up server error exit %d: %s", ec, reason)
		} else {
			logger.Errorf("Start up server error exit %d: failed", ec)
		}
		return ec
	} else {
		if reason != "" {
			log.Printf("Start up server error exit %d: %s", ec, reason)
		} else {
			log.Printf("Start up server error exit %d: failed", ec)
		}
		return ec
	}
}
