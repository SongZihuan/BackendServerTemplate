// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package globalmain

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func PreRun() (exitCode error) {
	var err error

	err = consoleutils.SetConsoleCPSafe(consoleutils.CodePageUTF8)
	if err != nil {
		return exitutils.InitFailedErrorForWin32ConsoleModule(err.Error())
	}

	err = logger.InitBaseLogger()
	if err != nil {
		return exitutils.InitFailedErrorForLoggerModule(err.Error())
	}

	return nil
}

func PostRun() {
	logger.CloseLogger()
	logger.Recover()
}
