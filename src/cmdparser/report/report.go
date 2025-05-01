// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package report

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/formatutils"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var CMD = &cobra.Command{
	Use:   "report",
	Short: "Print how to submit feedback",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, _ = printReport(os.Stdout)
		return nil
	},
}

func printReport(writer io.Writer) (int, error) {
	report := formatutils.FormatTextToWidth(global.Report, formatutils.NormalConsoleWidth)
	return fmt.Fprint(writer, report)
}
