// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/releaseinfo"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/touch"
	"github.com/spf13/cobra"
)

func init() {
	resource.Init()
}

var normalCommand = &cobra.Command{
	Use:   "normal",
	Short: "generate release info",
	RunE:  runNormal,
}

func runNormal(cmd *cobra.Command, args []string) error {
	exitreturn.SaveExitCode(normal())
	return nil
}

func normal() (exitcode int) {
	var err error

	genlog.GenLog("start to run")
	defer func() {
		genlog.GenLogf("run stop [code: %d]", exitcode)
	}()

	err = mod.InitGoModuleName() // 确保运行目录与 go.mod同级
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = touch.TouchReleaseFile()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = git.InitGitData()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = releaseinfo.WriteReleaseData(resource.Version)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
