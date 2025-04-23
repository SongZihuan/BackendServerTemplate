// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/formatutils"
	"github.com/spf13/cobra"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var short bool

var CMD = &cobra.Command{
	Use:   "version",
	Short: "Print the version and build info of this program",
	RunE: func(cmd *cobra.Command, args []string) error {
		if short {
			_, _ = printShortVersion(os.Stdout)
		} else {
			_, _ = printVersion(os.Stdout)
		}
		return nil
	},
}

func init() {
	CMD.Flags().BoolVarP(&short, "short", "s", false, "only show the version info")
}

func printVersion(writer io.Writer) (int, error) {
	res := new(strings.Builder)
	res.WriteString(fmt.Sprintf("Version: %s\n", global.Version))
	res.WriteString(fmt.Sprintf("Build Date (UTC): %s\n", global.BuildTime.In(global.UTCLocation).Format(time.DateTime)))
	if global.LocalLocation.String() != global.UTCLocation.String() {
		res.WriteString(fmt.Sprintf("Build Date (%s): %s\n", global.LocalLocation.String(), global.BuildTime.In(global.LocalLocation).Format(time.DateTime)))
	}
	res.WriteString(fmt.Sprintf("Compiler: %s\n", runtime.Version()))
	res.WriteString(fmt.Sprintf("OS: %s\n", runtime.GOOS))
	res.WriteString(fmt.Sprintf("Arch: %s\n", runtime.GOARCH))

	version := formatutils.FormatTextToWidth(res.String(), formatutils.NormalConsoleWidth)
	return fmt.Fprint(writer, version)
}

func printShortVersion(writer io.Writer) (int, error) {
	return fmt.Fprint(writer, global.SemanticVersioning) // 不需要(ln)换行
}
