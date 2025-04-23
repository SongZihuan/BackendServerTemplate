// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
)

var config *configInfo

func InitConfig(opt *ConfigOption) error {
	if config != nil {
		return fmt.Errorf("config already init")
	}

	_cfg, cfgErr := newConfig(opt)
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
		logger.Panic("config is not ready")
	}

	return config.GetData()
}

func Data() *ConfigData {
	if config == nil {
		logger.Panic("config is not ready")
	}

	return config.Data()
}

func OutputPath() string {
	if config == nil {
		logger.Panic("config is not ready")
	}

	return config.OutputPath()
}
