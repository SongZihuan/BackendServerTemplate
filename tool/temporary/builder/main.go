// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/packagelist"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/modutils"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var targetOS = runtime.GOOS
var targetArch = runtime.GOARCH
var targetPackage string

var packageMap map[string]string
var gomod string

var rootCommand = &cobra.Command{
	Use:   "builder",
	Short: "build the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runBuild(cmd, args, targetOS, targetArch, targetPackage)
	},
}

func init() {
	if envGOOS := envutils.GetSysEnv("GOOS"); envGOOS != "" {
		targetOS = envGOOS
	}

	if envArch := envutils.GetSysEnv("GOARCH"); envArch != "" {
		targetArch = envArch
	}

	rootCommand.Flags().StringVar(&targetOS, "os", targetOS, "target platform operating system")
	rootCommand.Flags().StringVar(&targetArch, "arch", targetArch, "target platform architecture")
	rootCommand.Flags().StringVar(&targetPackage, "target", targetPackage, "target name")
}

func Init() error {
	templog.InitTempLog("builder", os.Stdout)

	mod, err := modutils.GetGoModuleName()
	if err != nil {
		return err
	}

	pkgList, err := packagelist.GetPackageList(mod)
	if err != nil {
		return err
	}

	for pkg, path := range pkgList {
		templog.TempLogf("Package [%s]: %s", pkg, path)
	}

	packageMap = pkgList
	gomod = mod
	return nil
}

func main() {
	os.Exit(command())
}

func command() int {
	err := Init()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	err = rootCommand.Execute()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	return exitreturn.GetExitCode()
}
