// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package random

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/randomutils"
)

func WriteRandomData() error {
	return fileutils.Write(basefile.FileRandomData, randomutils.GenerateRandomString(40))
}
