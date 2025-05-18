// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/buildpath"
)

func release(goos string, goarch string, target string) error {
	releasePath := buildpath.TargetReleaseDir(goos, goarch, target)
	if releasePath == "" {
		return fmt.Errorf("release path not found: %s-%s-%s", goos, goarch, target)
	}

	outputPath := buildpath.TargetReleaseOutputFile(goos, goarch, target)
	if outputPath == "" {
		return fmt.Errorf("archive output path is empty: %s-%s-%s", goos, goarch, target)
	}

	err := createTarGZip(outputPath, releasePath)
	if err != nil {
		return err
	}

	return nil
}
