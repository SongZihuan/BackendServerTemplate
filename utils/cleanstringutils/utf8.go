// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

func CheckAndRemoveBOM(s string) string {
	// UTF-8 BOM 的字节序列为 0xEF, 0xBB, 0xBF
	bom := []byte{0xEF, 0xBB, 0xBF}

	// 将字符串转换为字节切片
	bytes := []byte(s)

	// 检查前三个字节是否是 BOM
	if len(bytes) >= 3 && bytes[0] == bom[0] && bytes[1] == bom[1] && bytes[2] == bom[2] {
		// 如果存在 BOM，则删除它
		return string(bytes[3:])
	}

	// 如果不存在 BOM，则返回原始字符串
	return s
}
