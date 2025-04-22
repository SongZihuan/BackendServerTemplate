// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package subcmd

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/check"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/license"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/report"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/version"
	"github.com/spf13/cobra"
)

func AddSubCMDOfRoot(cmd *cobra.Command) {
	cmd.AddCommand(version.CMD, license.CMD, report.CMD, check.CMD)
}
