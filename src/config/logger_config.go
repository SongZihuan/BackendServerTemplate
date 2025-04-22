// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/typeutils"
)

type LoggerConfig struct {
	LogLevel loglevel.LoggerLevel `json:"log-level" yaml:"log-level" mapstructure:"log-level"`
	LogTag   typeutils.StringBool `json:"log-tag" yaml:"log-tag" mapstructure:"log-tag"`

	HumanWarnWriter   LoggerWriterConfig `json:"human-warn-writer" yaml:"human-warn-writer" mapstructure:"human-warn-writer"`
	HumanErrWriter    LoggerWriterConfig `json:"human-error-writer" yaml:"human-error-writer" mapstructure:"human-error-writer"`
	MachineWarnWriter LoggerWriterConfig `json:"machine-warn-writer" yaml:"machine-warn-writer" mapstructure:"machine-warn-writer"`
	MachineErrWriter  LoggerWriterConfig `json:"machine-error-writer" yaml:"machine-error-writer" mapstructure:"machine-error-writer"`
}

func (d *LoggerConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	cfgErr := d.HumanWarnWriter.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.HumanErrWriter.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineWarnWriter.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineErrWriter.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *LoggerConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.LogLevel == "" {
		if c.data.GlobalConfig.IsRelease() {
			d.LogLevel = loglevel.LevelInfo
			d.LogTag.SetDefaultDisable()
		} else {
			d.LogLevel = loglevel.LevelDebug
			d.LogTag.SetDefaultEnable()
		}
	}

	cfgErr := d.HumanWarnWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.HumanErrWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineWarnWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineErrWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *LoggerConfig) check(c *configInfo) (err configerror.Error) {
	cfgErr := d.HumanWarnWriter.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.HumanErrWriter.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineWarnWriter.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineErrWriter.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *LoggerConfig) process(c *configInfo) (cfgErr configerror.Error) {
	err := logger.SetLevel(d.LogLevel)
	if err != nil {
		return configerror.NewErrorf("set log level error: %s", err.Error())
	}

	err = logger.SetLogTag(d.LogTag.IsEnable(false))
	if err != nil {
		return configerror.NewErrorf("set log tag error: %s", err.Error())
	}

	cfgErr = d.HumanWarnWriter.process(c, logger.SetHumanWarnWriter)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.HumanErrWriter.process(c, logger.SetHumanErrWriter)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineWarnWriter.process(c, logger.SetMachineWarnWriter)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.MachineErrWriter.process(c, logger.SetMachineErrWriter)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}
