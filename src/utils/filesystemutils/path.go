// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filesystemutils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func CleanFilePathAbs(path string) (string, error) {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		index := strings.Index(path, `:\`)
		pf := strings.ToUpper(path[:index])
		ph := path[index:]
		path = pf + ph
	}

	return path, nil
}
