// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package osutils

import (
	"os"
	"path/filepath"
)

var _args0 = ""

func init() {
	var err error
	if len(os.Args) > 0 {
		_args0, err = os.Executable()
		if err != nil {
			_args0 = os.Args[0]
		}
	}

	if _args0 == "" {
		panic("args was empty")
	}
}

func GetArgs0() string {
	return _args0
}

func GetArgs0Name() string {
	return filepath.Base(_args0)
}
