// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"github.com/spf13/cobra"
)

func runWinBuildAll(cmd *cobra.Command, args []string) error {
	exitreturn.SaveExitCode(allBuildAndRelease(cmd, args, runtimeutils.Windows))
	return nil
}

func runLinuxBuildAll(cmd *cobra.Command, args []string) error {
	exitreturn.SaveExitCode(allBuildAndRelease(cmd, args, runtimeutils.Linux))
	return nil
}

func allBuildAndRelease(cmd *cobra.Command, args []string, goos string) (exitcode int) {
	templog.TempLogf("check arch")
	archmap, ok := platformMap[goos]
	if !ok || len(archmap) == 0 {
		return exitreturn.ReturnError(fmt.Errorf("os %s not support", goos))
	}
	templog.TempLogf("check arch success")

	templog.TempLogf("environment preparation")
	err := environmentPreparation(goos)
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("environment preparation success")

	for goarch, _ := range archmap {
		for target, _ := range packageMap {
			err := func() error {
				templog.TempLogf("start to build %s-%s-%s", goos, goarch, target)
				err = build(goos, goarch, target)
				if err != nil {
					return err
				}
				templog.TempLogf("start to build %s-%s-%s success", goos, goarch, target)

				templog.TempLogf("start to release %s-%s-%s", goos, goarch, target)
				err = release(goos, goarch, target)
				if err != nil {
					return err
				}
				templog.TempLogf("start to release %s-%s-%s success", goos, goarch, target)

				return nil
			}()
			if err != nil {
				return exitreturn.ReturnError(err)
			}
		}
	}

	return exitreturn.ReturnSuccess()
}
