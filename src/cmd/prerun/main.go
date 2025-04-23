// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package prerun

// 必须明确导入 global 包 （虽然下面的import确实导入了 global 包，但此处重复写一次表示冗余，以免在某些情况下本包不适用 global 后，下方的导入被自动删除）
import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/stdutils"
)

import (
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func PreRun() (exitCode error) {
	var err error

	quiteMode := envutils.GetEnv("QUITE")
	if quiteMode != "" {
		err = stdutils.QuiteMode()
		if err != nil {
			stdutils.Recover()
			return exitutils.InitFailedForQuiteModeModule(err.Error())
		}
	}

	if global.UTCLocation == nil {
		return exitutils.InitFailedForTimeLocationModule("can not get utc location")
	}

	if global.LocalLocation == nil {
		return exitutils.InitFailedForTimeLocationModule("can not get local location")
	}

	err = consoleutils.SetConsoleCPSafe(consoleutils.CodePageUTF8)
	if err != nil {
		return exitutils.InitFailedForWin32ConsoleModule(err.Error())
	}

	err = logger.InitBaseLogger()
	if err != nil {
		return exitutils.InitFailedForLoggerModule(err.Error())
	}

	return nil
}

func PostRun() {
	logger.CloseLogger()
	logger.Recover()
	stdutils.CloseNullFile()
}
