// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/packagelist"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/platformlist"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"os"
	"path/filepath"
)

func main() {
	os.Exit(command())
}

func command() (exitcode int) {
	templog.InitTempLog("release makedir", os.Stdout)

	templog.TempLogf("get package list")
	pkgList, err := packagelist.GetPackageList("") // 无需传递 gomod , 因为实际不关心模块路径
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("get package list success")

	templog.TempLogf("get platform list")
	_, _, platformList, err := platformlist.GetPlatformList()
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("get platform list success")

	dest := "./RELEASE"

	for goos, goarchList := range platformList {
		for goarch, _ := range goarchList {
			for pkgName := range pkgList {
				dirPath := filepath.Join(dest, fmt.Sprintf("%s_%s_%s", goos, goarch, pkgName))

				if isExists, isFile := filesystemutils.IsExistsAndFile(dirPath); !isExists {
					err := os.MkdirAll(dirPath, 755)
					if err != nil {
						return exitreturn.ReturnError(err)
					}
					templog.TempLogf("create release dir (%s)", dirPath)
				} else if isFile {
					return exitreturn.ReturnError(fmt.Errorf("the release dir (%s) is file\n", dirPath))
				}

				cfgPath := filepath.Join(dirPath, "INSTALL.yaml")

				if isExists, isDir := filesystemutils.IsExistsAndDir(cfgPath); isExists && isDir {
					return exitreturn.ReturnError(fmt.Errorf("the release config file (%s) is dir\n", cfgPath))
				} else if !isExists {
					f, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
					if err != nil {
						return exitreturn.ReturnError(fmt.Errorf("create release install file (%s) failed: %s\n", cfgPath, err.Error()))
					}
					_ = f.Close()
					templog.TempLogf("create release config file (%s)", cfgPath)
				}
			}
		}
	}

	return exitreturn.ReturnSuccess()
}
