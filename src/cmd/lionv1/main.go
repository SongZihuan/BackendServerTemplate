// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	lionv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/lion/v1"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
)

func main() {
	defer logger.Recover()

	cmd := root.GetRootCMD("Multi-tasking background system",
		"A multi-task background system controlled by a controller to run multiple tasks concurrently",
		&lionv1.AutoReload,
		true,
		lionv1.MainV1)

	cmd.Flags().StringVarP(&lionv1.InputConfigFilePath, "config", "c", lionv1.InputConfigFilePath, "the file path of the configuration file")
	cmd.Flags().StringVarP(&lionv1.OutputConfigFilePath, "output-config", "o", lionv1.OutputConfigFilePath, "the file path of the output configuration file")

	exitutils.ExitQuite(cmd.Execute())
}
