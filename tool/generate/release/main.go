// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/spf13/cobra"
	"os"
)

// 冗余导入此包，该包包含必须导入的全部信息
import (
	_ "github.com/SongZihuan/BackendServerTemplate/global/pkgimport"
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

	rootCommand.AddCommand(normalCommand, specialCommand)
	err := rootCommand.Execute()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.GetExitCode()
}
