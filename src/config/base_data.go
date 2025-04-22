// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ConfigData struct {
	GlobalConfig `json:",inline" yaml:",inline" mapstructure:",squash"`
	Logger       LoggerConfig       `json:"logger" yaml:"logger" mapstructure:"logger"`
	Signal       SignalConfig       `json:"signal" yaml:"signal" mapstructure:"signal"`
	Win32Console Win32ConsoleConfig `json:"win32-console" yaml:"win32-console" mapstructure:"win32-console"`
	Server       ServerConfig       `json:"server" yaml:"server" mapstructure:"server"`
}

func (d *ConfigData) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	cfgErr := d.GlobalConfig.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Win32Console.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) setDefault(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Win32Console.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) check(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Win32Console.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) process(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Win32Console.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}
