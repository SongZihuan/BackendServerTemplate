// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lifecycle

// 必须明确导入 global 包 （虽然下面的import确实导入了 global 包，但此处重复写一次表示冗余，以免在某些情况下本包不适用 global 后，下方的导入被自动删除）
import (
	_ "github.com/SongZihuan/BackendServerTemplate/global/tzdata"
)
