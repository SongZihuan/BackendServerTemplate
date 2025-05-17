// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/releaseinfo"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/touch"
	"github.com/spf13/cobra"
)

var specialVersion string
var force bool

var specialCommand = &cobra.Command{
	Use:   "special",
	Short: "generate special release info",
	RunE:  runSpecial,
}

func init() {
	specialCommand.Flags().StringVar(&specialVersion, "version", "v0.0.0", "special version")
	specialCommand.Flags().BoolVarP(&force, "force", "f", false, "force set version")
}

func runSpecial(cmd *cobra.Command, args []string) error {
	exitreturn.SaveExitCode(special(specialVersion, force))
	return nil
}

func special(version string, force bool) (exitcode int) {
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

	err = releaseinfo.WriteSpecialReleaseData(version, force)
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.ReturnSuccess()
}
