// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package builddate

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"
	"time"
)

var buildTime = time.Now()

func WriteBuildDateData() error {
	return fileutils.Write(basefile.FileBuildDateTxt, fmt.Sprint(buildTime.Unix()))
}
