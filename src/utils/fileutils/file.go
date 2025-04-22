// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fileutils

import "os"

func IsFileOpen(file *os.File) bool {
	// 获取文件描述符
	fd := file.Fd()

	// 检查文件描述符是否有效
	if fd < 0 {
		return false
	}

	// 尝试对文件描述符进行简单的读/写操作以确认其状态
	fileInfo, err := file.Stat()
	if err != nil || fileInfo == nil {
		return false
	}

	return true
}
