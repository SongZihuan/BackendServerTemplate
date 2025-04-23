// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !systemtzdata

package global

import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/utils/timeutils/tzdata"
)

// 默认情况下加载 go 自带的时区包，除非使用 systemtzdata 明确使用系统的时区包
