// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lifecycle

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/stdutils"
	"sync"
)

var _postRunOnce sync.Once

func PostRun() {
	_postRunOnce.Do(func() {
		postRun()
	})
}

func postRun() {
	logger.CloseLogger()
	logger.Recover()
	stdutils.CloseNullFile()
}
