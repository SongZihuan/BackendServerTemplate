// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package osutils

import (
	"os"
	"path/filepath"
	"strings"
)

var _args0 = ""
var _args0Name = ""
var _args0NamePOSIX = ""

func init() {
	var err error
	_args0, err = os.Executable()
	if err != nil {
		if len(os.Args) > 0 {
			_args0 = os.Args[0]
		} else {
			panic(err)
		}
	}

	if _args0 == "" {
		panic("_args0 was empty")
	}

	_args0Name = filepath.Base(_args0)

	if _args0Name == "" {
		panic("_args0Name was empty")
	}

	_args0NamePOSIX = strings.TrimSuffix(_args0Name, ".exe")

	if _args0NamePOSIX == "" {
		panic("_args0NamePOSIX was empty")
	}
}

func GetArgs0() string {
	return _args0
}

func GetArgs0Name() string {
	return _args0Name
}

func GetArgs0NamePOSIX() string {
	return _args0NamePOSIX
}
