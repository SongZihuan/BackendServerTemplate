// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/giraffe"
	"github.com/SongZihuan/BackendServerTemplate/src/lifecycle"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

// 必须明确导入 lifecycle 包 （虽然下面的import确实导入了 prerun 包，但此处重复写一次表示冗余，以免在某些情况下本包不适用 prerun 后，下方的导入被自动删除）
import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/lifecycle"
)

func main() {
	err := lifecycle.PreRun()
	defer lifecycle.PostRun() // 此处defer放在err之前（因为RPreRun包含启动东西太多，虽然返回err，但不代表全部操作没成功，因此defer设置在这个位置，确保清理函数被调用。清理函数可以判断当前是否需要清理）
	if err != nil {
		exitutils.ErrorToExit(err).ClampAttribute().Exit()
		return // 不可达
	}

	exitutils.ErrorToExitQuite(giraffe.GetMainCommand().Execute()).ClampAttribute().Exit()
}
