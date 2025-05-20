// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basic"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/builddate"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/version"
	"os"
)

// 冗余导入此包，该包包含必须导入的全部信息
import (
	_ "github.com/SongZihuan/BackendServerTemplate/global/pkgimport"
)

func init() {
	resource.Init()
}

func main() {
	os.Exit(command())
}

func command() (exitcode int) {
	var err error

	genlog.InitGenLog("generate build", os.Stdout)

	genlog.GenLog("start to run")
	defer func() {
		genlog.GenLogf("run stop [code: %d]", exitcode)
	}()

	err = mod.InitGoModuleName() // 确定当前是否在 go.mod 同目录下运行
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = git.InitGitData()
	if err != nil && !errors.Is(err, git.ErrWithoutGit) {
		return exitreturn.ReturnError(err)
	}

	err = version.InitLongVersion(resource.Version)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = version.InitShortVersion(resource.Version)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = basic.WriteBasicData(resource.License, resource.Report)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = mod.WriteModuleNameData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = builddate.WriteBuildDateData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = git.WriteGitData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = version.WriteShortVersion(resource.Version)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = version.WriteLongVersion(resource.Version)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = builder.SetConfig(resource.BuildConfig)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = builder.SaveGlobalData(basefile.FileBuildDateGob)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = git.WriteGitIgnore()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
