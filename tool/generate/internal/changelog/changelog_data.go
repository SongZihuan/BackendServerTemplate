// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package changelog

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"os"
	"strings"
	"sync"
)

const FileChangelog = "./CHANGELOG.md"

var lastChangeLog = ""
var once sync.Once

func GetLastChangLog() string {
	once.Do(func() {
		lastChangeLog = getLastChangLog()
	})
	return lastChangeLog
}

func getLastChangLog() string {
	genlog.GenLogf("get %s data\n", FileChangelog)
	defer genlog.GenLogf("get %s data finish\n", FileChangelog)

	dat, err := os.ReadFile(FileChangelog)
	if err != nil {
		genlog.GenLogf("read file %s failed: %s\n", FileChangelog, err.Error())
		return ""
	}

	logSrc := strings.Split(cleanstringutils.GetString(string(dat)), "\n")

	res := new(strings.Builder)
	index := 0

	// 定位最新版本
FindVersionCycle:
	for ; ; index++ {
		if index >= len(logSrc) {
			genlog.GenLogf("read file %s failed: log title not found\n", FileChangelog)
			return ""
		}

		s := logSrc[index]
		if strings.HasPrefix(s, "## [") && !strings.HasPrefix(s, "## [未") {
			genlog.GenLogf("read file %s title [index: %d]: %s\n", FileChangelog, index, s)
			res.WriteString("\n\n---\n\n") // 前一个\n\n是用于与前者空开一行以上（使用双\n放在前者写完后没空行。后面\n\n双恐慌可以实现后者直接写内容就与前面保存空行。
			res.WriteString(fmt.Sprintf("### **%s** 更新内容\n", git.Version()))
			break FindVersionCycle
		}
	}

GetVersionLogCycle:
	for index++; ; index++ { // 初始化的 index++ 是为了从标题哪一行向下走一行。因为对于上面的循环（FindVersionCycle）来说，break 后置语句（index++）也就不会执行，若本循环开头不执行（index++）则会导致本循环读取的第一条数据和 FindVersionCycle 循环读取的第一条数据为同一行（标题行）
		if index >= len(logSrc) {
			break GetVersionLogCycle
		}

		s := logSrc[index]
		if strings.HasPrefix(s, "## [") {
			genlog.GenLogf("read file %s content end [index: %d]\n", FileChangelog, index)
			break GetVersionLogCycle
		}

		genlog.GenLogf("read file %s content [index: %d]: %s\n", FileChangelog, index, s)
		if strings.HasPrefix(s, "### ") {
			res.WriteString("#" + s + "\n")
		} else {
			res.WriteString(s + "\n")
		}
	}

	return res.String()
}
