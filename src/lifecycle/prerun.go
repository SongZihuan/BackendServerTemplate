// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lifecycle

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/buildinfo"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/runner"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/exitutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/stdutils"
	"sync"
)

// 冗余导入此包，该包包含必须导入的全部信息
import (
	_ "github.com/SongZihuan/BackendServerTemplate/global/pkgimport"
)

var _preRunOnce sync.Once
var _preError error

func PreRun(packageName string) (exitCode error) {
	_preRunOnce.Do(func() {
		_preError = preRun(packageName)
	})
	return _preError
}

func preRun(packageName string) (exitCode error) {
	var err error

	if packageName == "" {
		return exitutils.InitFailed("Global Data", "package is empty")
	} else if newPackageName := cleanstringutils.GetString(packageName); packageName != newPackageName {
		return exitutils.InitFailed("Global Data", fmt.Sprintf("package is invalid, use %s please", newPackageName))
	}

	err = runner.ReadGlobalData(buildinfo.BuildData, packageName)
	if err != nil {
		return exitutils.InitFailed("Global Data", err.Error())
	}

	err = rtdata.SetName(runner.GetConfigNamePointer(), runner.GetConfigAutoNamePointer(), packageName)
	if err != nil {
		return exitutils.InitFailed("Global Data", err.Error())
	}

	quiteMode := envutils.GetEnv(runner.GetConfigEnvPrefix(), "QUITE")
	if quiteMode != "" {
		err = stdutils.QuiteMode()
		if err != nil {
			stdutils.Recover()
			return exitutils.InitFailed("Quite Mode", err.Error())
		}
	}

	if rtdata.GetUTC() == nil {
		return exitutils.InitFailed("Time Location", "can not get utc location")
	}

	if rtdata.GetLocalLocation() == nil {
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
