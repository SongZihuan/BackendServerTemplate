// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
	"github.com/spf13/cobra"
)

func MainV1(cmd *cobra.Command, args []string, inputConfigFilePath string, outputConfigFilePath string) (exitCode error) {
	var err error

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: inputConfigFilePath,
		OutputFilePath: outputConfigFilePath,
		AutoReload:     false,
	})
	if err != nil {
		return exitutils.RunError(fmt.Sprintf("config file check failed: %s!", err.Error()))
	}

	outputPath := config.OutputPath()
	if outputPath != "" && filesystemutils.IsFile(outputPath) {
		logger.Warnf("config output ok: %s", outputPath)
	}

	return exitutils.SuccessExit("config file check ok!")
}
