// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package changelog

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/cleanstringutils"
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
	dat, err := os.ReadFile(FileChangelog)
	if err != nil {
		panic(err)
	}

	logSrc := strings.Split(cleanstringutils.GetString(string(dat)), "\n")

	res := new(strings.Builder)
	index := 0

	// 定位最新版本
FindVersionCycle:
	for {
		if index >= len(logSrc) {
			panic("Error CHANGELOG.md： can not find the log")
		}

		s := logSrc[index]
		if strings.HasPrefix(s, "## [") && !strings.HasPrefix(s, "## [未") {
			res.WriteString(s + "\n")
			break FindVersionCycle
		}
	}

GetVersionLogCycle:
	for {
		if index >= len(logSrc) {
			break GetVersionLogCycle
		}

		s := logSrc[index]
		if strings.HasPrefix(s, "## [") {
			break GetVersionLogCycle
		}

		res.WriteString(s + "\n")
	}

	return res.String()
}
