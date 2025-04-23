// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tzdata

import (
	_ "time/tzdata"
)

// 导入 Go 内置的时区数据（用于系统没有时区数据时，例如：Docker）
