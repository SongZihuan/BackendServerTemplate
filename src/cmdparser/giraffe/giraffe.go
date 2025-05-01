// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package giraffe

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/lion"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/tiger"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/spf13/cobra"
)

func GetMainCommand() *cobra.Command {
	cmd := GetCommand(global.Name)
	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD)
	return cmd
}

func GetCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           name,
		Short:         "Background server system",
		Long:          "Background server system, include tiger and lion mode",
		SilenceUsage:  false,
		SilenceErrors: false,
	}

	cmd.AddCommand(lion.GetCommand("lion"), tiger.GetCommand("tiger"))

	return cmd
}
