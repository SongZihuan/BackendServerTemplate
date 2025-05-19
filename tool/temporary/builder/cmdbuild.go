// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/spf13/cobra"
)

func runBuild(cmd *cobra.Command, args []string, goos string, goarch string, target string) error {
	exitreturn.SaveExitCode(buildAndRelease(cmd, args, goos, goarch, target))
	return nil
}

func buildAndRelease(cmd *cobra.Command, args []string, goos string, goarch string, target string) (exitcode int) {
	templog.TempLogf("environment preparation")
	err := environmentPreparation(goos)
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("environment preparation success")

	templog.TempLogf("start to build %s-%s-%s", goos, goarch, target)
	err = build(goos, goarch, target)
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("start to build %s-%s-%s success", goos, goarch, target)

	templog.TempLogf("start to release %s-%s-%s", goos, goarch, target)
	err = release(goos, goarch, target)
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("start to release %s-%s-%s success", goos, goarch, target)

	return exitreturn.ReturnSuccess()
}
