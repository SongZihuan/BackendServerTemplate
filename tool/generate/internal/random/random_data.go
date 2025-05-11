// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package random

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/randomutils"
)

func WriteRandomData() error {
	genlog.GenLog("write random number (length=40) data")
	defer genlog.GenLog("write random number (length=40) data finish")

	val := randomutils.GenerateRandomString(40, "")

	genlog.GenLogf("write %s to file %s\n", val, basefile.FileRandomData)
	return fileutils.Write(basefile.FileRandomData, val)
}
