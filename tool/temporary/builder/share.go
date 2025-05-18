// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/executils"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"os"
	"os/exec"
)

const GoModFile = "./go.mod"
const BuildConfigFile = "./BUILD.yaml"
const OutputDir = "./OUTPUT"

var ErrNotAdmin = fmt.Errorf("error not admin")

func targetReleaseDir(goos string, goarch string, target string) string {
	return fmt.Sprintf("./RELEASE/%s_%s_%s", goos, goarch, target)
}

func targetReleaseOutput(goos string, goarch string, target string) string {
	if goos == runtimeutils.Windows {
		return fmt.Sprintf("./RELEASE/%s_%s_%s/%s.exe", goos, goarch, target, target)
	}
	return fmt.Sprintf("./RELEASE/%s_%s_%s/%s", goos, goarch, target, target)
}

func adminTargetReleaseOutput(goos string, goarch string, target string) string {
	if goos != runtimeutils.Windows {
		return ""
	}
	return fmt.Sprintf("./RELEASE/%s_%s_%s/%s-admin.exe", goos, goarch, target, target)
}

func targetBuildFileName(goos string, goarch string) string {
	return fmt.Sprintf("./BUILD/BUILD.%s.%s.yaml", goos, goarch)
}

func adminTargetBuildFileName(goos string, goarch string) string {
	if goos != runtimeutils.Windows {
		return ""
	}

	return fmt.Sprintf("./BUILD/BUILD.%s.%s.admin.yaml", goos, goarch)
}

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

	if !filesystemutils.IsFile(targetBuildFileName(goos, goarch)) {
		return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
	}

	if goos != runtimeutils.Windows {
		if !filesystemutils.IsFile(adminTargetBuildFileName(goos, goarch)) {
			return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
		}
	}

	if !filesystemutils.IsDir(targetReleaseDir(goos, goarch, target)) {
		return fmt.Errorf("os [%s] or arch [%s] is no support", goos, goarch)
	}

	return nil
}

func parameterAdminCheck(goos string, goarch string) error {
	if goos != runtimeutils.Windows {
		return ErrNotAdmin
	}

	if !filesystemutils.IsFile(adminTargetBuildFileName(goos, goarch)) {
		return ErrNotAdmin
	}

	return nil
}

func environmentPreparation(goos string, goarch string, target string) error {
	if !filesystemutils.IsFile(GoModFile) {
		return fmt.Errorf("please run in the root path (which has the file go.mod)")
	}

	if filesystemutils.IsFile(BuildConfigFile) {
		err := os.Remove(BuildConfigFile)
		if err != nil {
			return err
		}
	}

	if filesystemutils.IsFile(OutputDir) {
		return fmt.Errorf("the %s is not a dir", OutputDir)
	}

	err := os.MkdirAll(OutputDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func copyBuildConfigFile(goos string, goarch string) error {
	_, err := fileutils.Copy(BuildConfigFile, targetBuildFileName(goos, goarch))
	return err
}

func copyAdminBuildConfigFile(goos string, goarch string) error {
	if goos != runtimeutils.Windows {
		return fmt.Errorf("os [%s] is not support", goos)
	}

	_, err := fileutils.Copy(BuildConfigFile, adminTargetBuildFileName(goos, goarch))
	return err
}

func build(goos string, goarch string, target string) (exitcode int) {
	templog.TempLogf("start to build %s-%s-%s", goos, goarch, target)
	defer func() {
		templog.TempLogf("run stop [code: %d]", exitcode)
	}()

	templog.TempLogf("build base program")
	err := buildBase(goos, goarch, target)
	if err != nil {
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("build base program success")

	templog.TempLogf("build program with admin")
	err = buildAdmin(goos, goarch, target)
	if err != nil && !errors.Is(err, ErrNotAdmin) {
		templog.TempLogf("build program with admin not support")
		return exitreturn.ReturnError(err)
	}
	templog.TempLogf("build program with admin success")

	return exitreturn.ReturnSuccess()
}

func buildBase(goos string, goarch string, target string) error {
	templog.TempLogf("check the parameter")
	err := parameterCheck(goos, goarch, target)
	if err != nil {
		return err
	}
	templog.TempLogf("check the parameter success")

	templog.TempLogf("environment preparation")
	err = environmentPreparation(goos, goarch, target)
	if err != nil {
		return err
	}
	templog.TempLogf("environment preparation success")

	templog.TempLogf("copy BUILD.yaml")
	err = copyBuildConfigFile(goos, goarch)
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

	output := targetReleaseOutput(goos, goarch, target)

	templog.TempLogf("go build: %s-%s-%s -> %s", goos, goarch, target, output)
	err = goBuild(goos, goarch, target, output)
	if err != nil {
		return err
	}
	templog.TempLogf("go build: %s-%s-%s -> %s success", goos, goarch, target, output)

	return nil
}

func buildAdmin(goos string, goarch string, target string) error {
	templog.TempLogf("check the admin parameter")
	err := parameterAdminCheck(goos, goarch)
	if err != nil {
		return err
	}
	templog.TempLogf("check the admin parameter success")

	templog.TempLogf("environment preparation")
	err = environmentPreparation(goos, goarch, target)
	if err != nil {
		return err
	}
	templog.TempLogf("environment preparation success")

	templog.TempLogf("copy BUILD.yaml")
	err = copyAdminBuildConfigFile(goos, goarch)
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

	output := adminTargetReleaseOutput(goos, goarch, target)

	templog.TempLogf("go build: %s-%s-%s -> %s", goos, goarch, target, output)
	err = goBuild(goos, goarch, target, output)
	if err != nil {
		return err
	}
	templog.TempLogf("go build: %s-%s-%s -> %s success", goos, goarch, target, output)

	return nil
}

func goGenerate() error {
	_, err := executils.Run("go", "generate", "./...")
	if err != nil {
		return err
	}

	return nil
}

func goBuild(goos string, goarch string, target string, output string) error {
	packagePath, ok := packageMap[target]
	if !ok {
		return fmt.Errorf("target [%s] is invalid", target)
	}

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		output,
		"-trimpath",
		`-ldflags`,
		`-s -w -extldflags "-static"`,
		`-gcflags`,
		`all=-l=4`,
		packagePath,
	)

	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env,
		"GOOS="+goos,
		"GOARCH="+goarch,
		"CGO=1",
	)

	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
