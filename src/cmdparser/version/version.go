// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/runner"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/utils/formatutils"
	"github.com/spf13/cobra"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var short bool
var version bool

var CMD = &cobra.Command{
	Use:   "version",
	Short: "Print the version and build info of this program",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version && short {
			_, _ = printShortVersion(os.Stdout)
		} else if !version && short {
			_, _ = printShortVersionInfo(os.Stdout)
		} else if version && !short {
			_, _ = printLongVersion(os.Stdout)
		} else {
			_, _ = printLongVersionInfo(os.Stdout)
		}
		return nil
	},
}

func init() {
	CMD.Flags().BoolVarP(&version, "version", "v", false, "only show the version")
	CMD.Flags().BoolVarP(&short, "short", "s", false, "show the short version")
}

func printLongVersionInfo(writer io.Writer) (int, error) {
	res := new(strings.Builder)
	res.WriteString(fmt.Sprintf("Version: %s\n", runner.GetLongVersion()))
	res.WriteString(fmt.Sprintf("Build Date (UTC): %s\n", runner.GetBuildDate().In(rtdata.GetUTC()).Format(time.DateTime)))
	if rtdata.GetLocalLocation().String() != rtdata.GetUTC().String() {
		res.WriteString(fmt.Sprintf("Build Date (%s): %s\n", rtdata.GetLocalLocation().String(), runner.GetBuildDate().In(rtdata.GetLocalLocation()).Format(time.DateTime)))
	}
	res.WriteString(fmt.Sprintf("Compiler: %s\n", runtime.Version()))
	res.WriteString(fmt.Sprintf("OS: %s\n", runtime.GOOS))
	res.WriteString(fmt.Sprintf("Arch: %s\n", runtime.GOARCH))

	return fmt.Fprint(writer, formatutils.FormatTextToWidth(res.String(), formatutils.NormalConsoleWidth))
}

func printShortVersionInfo(writer io.Writer) (int, error) {
	res := new(strings.Builder)
	res.WriteString(fmt.Sprintf("Version: %s\n", runner.GetShortVersion()))
	res.WriteString(fmt.Sprintf("Build Date (UTC): %s\n", runner.GetBuildDate().In(rtdata.GetUTC()).Format(time.DateTime)))
	if rtdata.GetLocalLocation().String() != rtdata.GetUTC().String() {
		res.WriteString(fmt.Sprintf("Build Date (%s): %s\n", rtdata.GetLocalLocation().String(), runner.GetBuildDate().In(rtdata.GetLocalLocation()).Format(time.DateTime)))
	}
	res.WriteString(fmt.Sprintf("Compiler: %s\n", runtime.Version()))
	res.WriteString(fmt.Sprintf("OS: %s\n", runtime.GOOS))
	res.WriteString(fmt.Sprintf("Arch: %s\n", runtime.GOARCH))

	return fmt.Fprint(writer, formatutils.FormatTextToWidth(res.String(), formatutils.NormalConsoleWidth))
}

func printLongVersion(writer io.Writer) (int, error) {
	return fmt.Fprint(writer, runner.GetLongSemanticVersion()) // 不需要(ln)换行
}

func printShortVersion(writer io.Writer) (int, error) {
	return fmt.Fprint(writer, runner.GetShortSemanticVersion()) // 不需要(ln)换行
}
