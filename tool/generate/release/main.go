// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/global"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "release-generate",
	Short: "generate release info",
	RunE:  runNormal,
}

func main() {
	os.Exit(command())
}

func command() int {
	genlog.InitGenLog("generate release", os.Stdout)

	err := global.GenerateReleaseInit()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	rootCommand.AddCommand(normalCommand, specialCommand)
	err = rootCommand.Execute()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.GetExitCode()
}
