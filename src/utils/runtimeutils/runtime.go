// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtimeutils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetCallingFunctionInfo(skip int) (string, string, string, int) {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", "", "", 0
	}

	var funcName string
	tmp := runtime.FuncForPC(pc).Name()
	tmpLst := strings.Split(tmp, "/")
	if len(tmpLst) == 0 {
		funcName = tmp
	} else {
		funcName = tmpLst[len(tmpLst)-1]
	}

	return funcName, file, filepath.Base(file), line
}
