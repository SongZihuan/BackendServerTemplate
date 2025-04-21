// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package license

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/formatutils"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var CMD = &cobra.Command{
	Use:   "license",
	Short: "Print the license of this project",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, _ = printLicense(os.Stdout)
		return nil
	},
}

func printLicense(writer io.Writer) (int, error) {
	license := formatutils.FormatTextToWidth(global.License, formatutils.NormalConsoleWidth)
	return fmt.Fprint(writer, license)
}
