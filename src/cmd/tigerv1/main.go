// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	tigerv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/tiger/v1"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func main() {
	cmd := root.GetRootCMD("Single-tasking background system",
		"A single-task background system that runs a single task directly without using a controller",
		tigerv1.MainV1)

	cmd.Flags().StringVarP(&tigerv1.InputConfigFilePath, "config", "c", tigerv1.InputConfigFilePath, "the file path of the configuration file")
	cmd.Flags().StringVarP(&tigerv1.InputConfigFilePath, "output-config", "o", tigerv1.InputConfigFilePath, "the file path of the output configuration file")

	exitutils.Exit(cmd.Execute())
}
