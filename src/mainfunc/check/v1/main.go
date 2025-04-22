// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/spf13/cobra"
)

var InputConfigFilePath string = "config.yaml"
var OutputConfigFilePath string = ""

func MainV1(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	configProvider, err := configparser.NewProvider(InputConfigFilePath, &configparser.NewProviderOption{
		AutoReload: false,
	})
	if err != nil {
		return exitutils.SuccessExitSimple(fmt.Sprintf("Error: config file check failed: %s!", err.Error()))
	}

	err = config.InitConfig(&config.ConfigOption{
		ConfigFilePath: InputConfigFilePath,
		OutputFilePath: OutputConfigFilePath,
		Provider:       configProvider,
	})
	if err != nil {
		return exitutils.SuccessExitSimple(fmt.Sprintf("Error: config file check failed: %s!", err.Error()))
	}

	return exitutils.SuccessExitSimple("Info: config file check ok!")
}
