// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/packagelist"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"os"
	"path/filepath"
	"strings"
)

// 必须运行在项目跟目录
func main() {
	os.Exit(command())
}

func command() (exitcode int) {
	templog.InitTempLog("release makedir", os.Stdout)
	pkgList, err := packagelist.GetPackageList()
	if err != nil {
		return exitreturn.ReturnError(err)
	}

	src := "./BUILD"
	dest := "./RELEASE"

	count := 0
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			templog.TempLogf("walk error: %s, skip %s\n", err.Error(), path)
			return nil
		}

		name := info.Name()

		if info.IsDir() {
			if count == 0 && name == "BUILD" {
				return nil
			} else {
				return filepath.SkipDir
			}
		}

		count++

		if !strings.HasPrefix(name, "BUILD.") {
			templog.TempLogf("walk error: not prefix `BUILD.` , skip %s\n", path)
			return nil
		} else if !strings.HasSuffix(name, ".yaml") {
			templog.TempLogf("walk error: not suffix `.yaml` , skip %s\n", path)
			return nil
		}

		osarch := strings.Split(strings.TrimPrefix(strings.TrimSuffix(name, ".yaml"), "BUILD."), ".")

		if len(osarch) != 2 && len(osarch) != 3 {
			templog.TempLogf("walk error: unknown file name, skip %s\n", path)
			return nil
		}

		maparch, ok := runtimeutils.ServerOSArch[osarch[0]]
		if !ok {
			templog.TempLogf("walk error: bad os %s , skip %s\n", osarch[0], path)
			return nil
		}

		yes, ok := maparch[osarch[1]]
		if !yes || !ok {
			templog.TempLogf("walk error: bad arch %s , skip %s\n", osarch[1], path)
			return nil
		}

		if len(osarch) == 3 {
			if osarch[0] != runtimeutils.Windows {
				templog.TempLogf("walk error: bad os %s with admin mode , skip %s\n", osarch[0], path)
				return nil
			}

			if osarch[2] != "admin" {
				templog.TempLogf("walk error: unknown file name, skip %s\n", path)
				return nil
			}
		}

		for _, pkg := range pkgList {
			dirPath := filepath.Join(dest, fmt.Sprintf("%s_%s_%s", osarch[0], osarch[1], pkg))

			err := os.MkdirAll(dirPath, 755)
			if err != nil {
				templog.TempLogf("create release dir (%s) failed: %s\n", dirPath, err.Error())
				continue
			}

			cfgPath := filepath.Join(dirPath, "INSTALL.yaml")
			isExists, isFile := filesystemutils.IsExistsAndFile(cfgPath)

			if isExists && isFile {
				err := os.RemoveAll(cfgPath)
				if err != nil {
					templog.TempLogf("create release install file (%s) failed: remove dir failed: %s\n", cfgPath, err.Error())
					continue
				}
				isExists = false
			}

			if !isExists {
				f, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
				if err != nil {
					templog.TempLogf("create release install file (%s) failed: %s\n", cfgPath, err.Error())
					continue
				}
				_ = f.Close()
			}
		}

		return nil
	})
	if err != nil {
		return exitreturn.ReturnError(fmt.Errorf("walk %s error: %v\n", src, err))
	}

	return exitreturn.ReturnSuccess()
}
