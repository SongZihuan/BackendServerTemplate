// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/patch"
	"os"
)

func main() {
	os.Exit(command())
}

func command() int {
	var err error

	genlog.InitGenLog("generate patch", nil)

	err = patch.InitPatchData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = patch.CreatePatch()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
