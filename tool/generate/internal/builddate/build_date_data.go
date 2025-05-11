// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package builddate

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"time"
)

var buildTime = time.Now()

func WriteBuildDateData() error {
	val := fmt.Sprint(buildTime.Unix())

	genlog.GenLogf("write build data: %s UTC (timestaamp: %s)\n", buildTime.UTC().Format(time.DateTime), val)
	defer genlog.GenLog("write build data finish")

	genlog.GenLogf("write %s to file %s\n", val, basefile.FileBuildDateTxt)
	return fileutils.Write(basefile.FileBuildDateTxt, val)
}
