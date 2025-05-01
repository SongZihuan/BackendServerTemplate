// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lifecycle

import (
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/exitutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/stdutils"
	"sync"
)

var _preRunOnce sync.Once
var _preError error

func PreRun() (exitCode error) {
	_preRunOnce.Do(func() {
		_preError = preRun()
	})
	return _preError
}

func preRun() (exitCode error) {
	var err error

	quiteMode := envutils.GetEnv(global.EnvPrefix, "QUITE")
	if quiteMode != "" {
		err = stdutils.QuiteMode()
		if err != nil {
			stdutils.Recover()
			return exitutils.InitFailed("Quite Mode", err.Error())
		}
	}

	if global.UTCLocation == nil {
		return exitutils.InitFailed("Time Location", "can not get utc location")
	}

	if global.LocalLocation == nil {
		return exitutils.InitFailed("Time Location", "can not get local location")
	}

	err = consoleutils.SetConsoleCPSafe(consoleutils.CodePageUTF8)
	if err != nil {
		return exitutils.InitFailed("Win32 Console API", err.Error())
	}

	err = logger.InitBaseLogger()
	if err != nil {
		return exitutils.InitFailed("Logger", err.Error())
	}

	err = stdutils.OpenNullFile()
	if err != nil {
		return exitutils.InitFailed("File dev/null", err.Error())
	}

	return nil
}
