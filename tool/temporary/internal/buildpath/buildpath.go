// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package buildpath

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
)

const GoModFile = "./go.mod"
const BuildConfigFile = "./BUILD.yaml"
const OutputDir = "./OUTPUT"

func TargetReleaseOutputFile(goos string, goarch string, target string) string {
	return fmt.Sprintf("./OUTPUT/%s_%s_%s.tar.gz", goos, goarch, target)
}

func AdminTargetReleaseOutputFile(goos string, goarch string, target string) string {
	return TargetReleaseOutputFile(goos, goarch, target)
}

func TargetReleaseDir(goos string, goarch string, target string) string {
	return fmt.Sprintf("./RELEASE/%s_%s_%s", goos, goarch, target)
}

func AdminTargetReleaseDir(goos string, goarch string, target string) string {
	return TargetReleaseDir(goos, goarch, target)
}

func TargetReleaseOutput(goos string, goarch string, target string) string {
	if goos == runtimeutils.Windows {
		return fmt.Sprintf("./RELEASE/%s_%s_%s/%s.exe", goos, goarch, target, target)
	}
	return fmt.Sprintf("./RELEASE/%s_%s_%s/%s", goos, goarch, target, target)
}

func AdminTargetReleaseOutput(goos string, goarch string, target string) string {
	if goos != runtimeutils.Windows {
		return ""
	}
	return fmt.Sprintf("./RELEASE/%s_%s_%s/%s-admin.exe", goos, goarch, target, target)
}

func TargetBuildFileName(goos string, goarch string) string {
	return fmt.Sprintf("./BUILD/BUILD.%s.%s.yaml", goos, goarch)
}

func AdminTargetBuildFileName(goos string, goarch string) string {
	if goos != runtimeutils.Windows {
		return ""
	}

	return fmt.Sprintf("./BUILD/BUILD.%s.%s.admin.yaml", goos, goarch)
}
