// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package check

import (
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/globalmain"
	checkv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/check/v1"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:     "check",
	Aliases: []string{"config"},
	Short:   "Check the correctness of the configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		err := globalmain.PreRun(false)
		if err != nil {
			return err
		}

		return checkv1.MainV1(cmd, args)
	},
}

func init() {
	CMD.Flags().StringVarP(&checkv1.InputConfigFilePath, "config", "c", checkv1.InputConfigFilePath, "the file path of the configuration file")
	CMD.Flags().StringVarP(&checkv1.OutputConfigFilePath, "output-config", "o", checkv1.OutputConfigFilePath, "the file path of the output configuration file")
}
