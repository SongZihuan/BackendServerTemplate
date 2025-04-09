// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resource

import (
	"regexp"
	"strings"
)

const semVerRegexStr = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`

var semVerRegex = regexp.MustCompile(semVerRegexStr)

// IsSemanticVersion checks if the given string is a valid semantic version.
func utilsIsSemanticVersion(version string) bool {
	return semVerRegex.MatchString(version)
}

func utilsClenFileDataMoreLine(data string) (res string) {
	res = utilsCheckAndRemoveBOM(data)
	res = strings.Replace(res, "\r", "", -1)
	return res
}

func utilsCheckAndRemoveBOM(s string) string {
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
