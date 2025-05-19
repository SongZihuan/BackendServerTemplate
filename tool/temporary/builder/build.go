// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/buildpath"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
)

// 编译（虽然可以通过协程提高速度，但程序本身依赖文件系统上的数据，因此无法并发编译，因为无法在同一时刻为程序准备多个不同的编译环境）
func build(goos string, goarch string, target string) error {
	templog.TempLogf("start to build %s-%s-%s", goos, goarch, target)
	defer func() {
		templog.TempLogf("build %s-%s-%s finish", goos, goarch, target)
	}()

	var hasAdmin = false

	templog.TempLogf("check the parameter")
	err := parameterCheck(goos, goarch, target)
	if err != nil {
		return err
	}
	templog.TempLogf("check the parameter success")

	templog.TempLogf("check the admin parameter")
	err = parameterAdminCheck(goos, goarch, target)
	if err == nil {
		hasAdmin = true
	} else if errors.Is(err, ErrNotAdmin) {
		hasAdmin = false
	} else {
		return err
	}
	templog.TempLogf("check the admin parameter success")

	templog.TempLogf("build base program")
	err = buildBase(goos, goarch, target)
	if err != nil {
		return err
	}
	templog.TempLogf("build base program success")

	if hasAdmin {
		templog.TempLogf("build program with admin")
		err = buildAdmin(goos, goarch, target)
		if err != nil && !errors.Is(err, ErrNotAdmin) {
			templog.TempLogf("build program with admin not support")
			return err
		}
		templog.TempLogf("build program with admin success")
	}

	return nil
}

func buildBase(goos string, goarch string, target string) error {
	templog.TempLogf("copy BUILD.yaml")
	err := copyBuildConfigFile(goos, goarch)
	if err != nil {
		return err
	}
	templog.TempLogf("copy BUILD.yaml success")

	templog.TempLogf("go generate")
	err = goGenerate()
	if err != nil {
		return err
	}
	templog.TempLogf("go generate success")

	output := buildpath.TargetReleaseOutput(goos, goarch, target)

	templog.TempLogf("go build: %s-%s-%s -> %s", goos, goarch, target, output)
	err = goBuild(goos, goarch, target, output)
	if err != nil {
		return err
	}
	templog.TempLogf("go build: %s-%s-%s -> %s success", goos, goarch, target, output)

	return nil
}

func buildAdmin(goos string, goarch string, target string) error {
	if goos != runtimeutils.Windows {
		return fmt.Errorf("os [%s] no support", goos)
	}

	templog.TempLogf("copy BUILD.yaml")
	err := copyAdminBuildConfigFile(goos, goarch)
	if err != nil {
		return err
	}
	templog.TempLogf("copy BUILD.yaml success")

	templog.TempLogf("go generate")
	err = goGenerate()
	if err != nil {
		return err
	}
	templog.TempLogf("go generate success")

	output := buildpath.AdminTargetReleaseOutput(goos, goarch, target)

	templog.TempLogf("go build: %s-%s-%s -> %s", goos, goarch, target, output)
	err = goBuild(goos, goarch, target, output)
	if err != nil {
		return err
	}
	templog.TempLogf("go build: %s-%s-%s -> %s success", goos, goarch, target, output)

	if mtptogram == "" {
		templog.TempLogf("set windows manifest by copy: %s", output)
		err = winCopyManifest(output, buildpath.WinAdminManifestFile)
	} else {
		templog.TempLogf("set windows manifest by mt.exe: %s", output)
		err = winMTManifest(output, buildpath.WinAdminManifestFile)
	}
	if err != nil {
		return err
	}
	templog.TempLogf("set windows manifest: %s success", output)

	return nil
}
