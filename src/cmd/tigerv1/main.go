// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/subcmd"
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	tigerv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/tiger/v1"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func main() {
	defer logger.Recover()

	cmd := root.GetRootCMD("Single-tasking background system",
		"A single-task background system that runs a single task directly without using a controller",
		&tigerv1.AutoReload,
		true,
		tigerv1.MainV1)

	subcmd.AddSubCMDOfRoot(cmd)
	cmd.Flags().StringVarP(&tigerv1.InputConfigFilePath, "config", "c", tigerv1.InputConfigFilePath, "the file path of the configuration file")
	cmd.Flags().StringVarP(&tigerv1.OutputConfigFilePath, "output-config", "o", tigerv1.OutputConfigFilePath, "the file path of the output configuration file")

	exitutils.ExitQuite(cmd.Execute())
}
