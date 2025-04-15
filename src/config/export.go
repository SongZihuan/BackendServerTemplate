// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"os"
	"path"
	"strings"
)

var config *configInfo

type ConfigOption struct {
	ConfigFilePath string
	OutputFilePath string
	Provider       configparser.ConfigParserProvider
}

func (opt *ConfigOption) setDefault() error {
	if opt.ConfigFilePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			logger.Errorf("can not get work directory: %s", err.Error())
			return err
		}

		opt.ConfigFilePath = path.Join(wd, "config.yaml")
		opt.Provider = configparser.NewYamlProvider()
	}

	if opt.Provider == nil {
		if strings.HasSuffix(opt.ConfigFilePath, ".yaml") || strings.HasSuffix(opt.ConfigFilePath, ".yml") {
			opt.Provider = configparser.NewYamlProvider()
		} else if strings.HasSuffix(opt.ConfigFilePath, ".json") {
			opt.Provider = configparser.NewJsonProvider()
		} else {
			opt.Provider = configparser.NewYamlProvider()
		}
	}

	return nil
}

func InitConfig(opt *ConfigOption) error {
	if config != nil {
		return fmt.Errorf("config already init")
	}

	if opt == nil {
		opt = new(ConfigOption)
	}

	err := opt.setDefault()
	if err != nil {
		return err
	}

	_cfg, cfgErr := newConfig(opt.ConfigFilePath, opt.OutputFilePath, opt.Provider)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = _cfg.init()
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	config = _cfg
	return nil
}

func GetData() (*ConfigData, configerror.Error) {
	if config == nil {
		panic("config is not ready")
	}

	return config.GetData()
}

func Data() *ConfigData {
	if config == nil {
		panic("config is not ready")
	}

	return config.Data()
}

func Output(filePath string) error {
	if config == nil {
		panic("config is not ready")
	}

	return config.Output(filePath)
}
