// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"log"
)

func ErrorToExit(err error) ExitCode {
	var ec ExitCode
	if err == nil {
		return exitCodeDefaultSuccess
	} else if errors.As(err, &ec) {
		return ec
	} else {
		if logger.IsReady() {
			logger.Errorf("Exit %d: %s", ec, err.Error())
		} else {
			log.Printf("Exit %d: %s\n", ec, err.Error())
		}
		return exitCodeDefaultError
	}
}

func ErrorToExitQuite(err error) ExitCode {
	var ec ExitCode
	if err == nil {
		return exitCodeDefaultSuccess
	} else if errors.As(err, &ec) {
		return ec
	} else {
		return exitCodeDefaultError
	}
}
