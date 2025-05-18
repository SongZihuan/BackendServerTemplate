// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/exitreturn"
	"github.com/spf13/cobra"
)

func runBuild(cmd *cobra.Command, args []string, goos string, goarch string, target string) error {
	exitreturn.SaveExitCode(build(goos, goarch, target))
	return nil
}
