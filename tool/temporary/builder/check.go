// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/buildpath"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"os"
)

var ErrNotAdmin = fmt.Errorf("error not admin")

func parameterCheck(goos string, goarch string, target string) error {
	if goos == "" {
		return fmt.Errorf("OS can not be empty")
	}

	if goarch == "" {
		return fmt.Errorf("arch can not be empty")
	}

	if target == "" {
		return fmt.Errorf("target must be set")
	}

	if yes, ok := runtimeutils.ServerOS[goos]; !yes || !ok {
		return fmt.Errorf("os [%s] is error", goos)
	}

	if yes, ok := runtimeutils.ServerArch[goarch]; !yes || !ok {
		return fmt.Errorf("arch [%s] is error", goarch)
	}

	archMap, ok := runtimeutils.ServerOSArch[goos]
	if archMap == nil || !ok {
		return fmt.Errorf("os [%s] is error", goos)
	}

	if yes, ok := archMap[goarch]; !yes || !ok {
		return fmt.Errorf("arch [%s] is error", goarch)
	}

	if _, ok := packageMap[target]; !ok {
		return fmt.Errorf("target [%s] is invalid", target)
	}

	if !filesystemutils.IsFile(buildpath.TargetBuildFileName(goos, goarch)) {
		return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
	}

	if goos != runtimeutils.Windows {
		if !filesystemutils.IsFile(buildpath.AdminTargetBuildFileName(goos, goarch)) {
			return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
		}
	}

	if !filesystemutils.IsDir(buildpath.TargetReleaseDir(goos, goarch, target)) {
		return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
	}

	return nil
}

func parameterAdminCheck(goos string, goarch string, target string) error {
	if goos != runtimeutils.Windows {
		return ErrNotAdmin
	}

	if !filesystemutils.IsFile(buildpath.AdminTargetBuildFileName(goos, goarch)) {
		return ErrNotAdmin
	}

	if !filesystemutils.IsDir(buildpath.AdminTargetReleaseDir(goos, goarch, target)) {
		return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
	}

	return nil
}

func environmentPreparation(goos string, goarch string, target string) error {
	if !filesystemutils.IsFile(buildpath.GoModFile) {
		return fmt.Errorf("please run in the root path (which has the file go.mod)")
	}

	if filesystemutils.IsFile(buildpath.BuildConfigFile) {
		err := os.Remove(buildpath.BuildConfigFile)
		if err != nil {
			return err
		}
	}

	if filesystemutils.IsFile(buildpath.OutputDir) {
		return fmt.Errorf("the %s is not a dir", buildpath.OutputDir)
	}

	err := os.MkdirAll(buildpath.OutputDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func copyBuildConfigFile(goos string, goarch string) error {
	_, err := fileutils.Copy(buildpath.BuildConfigFile, buildpath.TargetBuildFileName(goos, goarch))
	return err
}

func copyAdminBuildConfigFile(goos string, goarch string) error {
	if goos != runtimeutils.Windows {
		return fmt.Errorf("os [%s] is not support", goos)
	}

	_, err := fileutils.Copy(buildpath.BuildConfigFile, buildpath.AdminTargetBuildFileName(goos, goarch))
	return err
}
