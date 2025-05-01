// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filesystemutils

import "os"

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !s.IsDir()
}

func IsExistsAndFile(path string) (exists, isFile bool) {
	s, err := os.Stat(path)
	if err != nil {
		return false, false
	}

	return true, !s.IsDir()
}
