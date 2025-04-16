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
	LogLevel loglevel.LoggerLevel `json:"log-level" yaml:"log-level"`
	LogTag   typeutils.StringBool `json:"log-tag" yaml:"log-tag"`

	WarnWriter LoggerWriterConfig `json:"warn-writer" yaml:"warn-writer"`
	ErrWriter  LoggerWriterConfig `json:"error-writer" yaml:"error-writer"`
}

func (d *LoggerConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	cfgErr := d.WarnWriter.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.ErrWriter.init(filePath, provider)
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

	cfgErr := d.WarnWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.ErrWriter.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *LoggerConfig) check(c *configInfo) (err configerror.Error) {
	cfgErr := d.WarnWriter.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.ErrWriter.check(c)
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

	cfgErr = d.WarnWriter.process(c, logger.SetWarnWriter)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.ErrWriter.process(c, logger.SetErrWriter)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}
