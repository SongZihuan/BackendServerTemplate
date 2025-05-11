// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/releaseinfo"
	"os"
)

func main() {
	os.Exit(command())
}

func command() (exitcode int) {
	var err error

	genlog.InitGenLog("generate release", nil)

	genlog.GenLog("start to run")
	defer func() {
		genlog.GenLogf("run stop [code: %d]", exitcode)
	}()

	_, err = mod.GetGoModuleName() // 提前一步帕胺的
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = basefile.TouchReleaseFile()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = git.InitGitData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = releaseinfo.WriteReleaseData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
