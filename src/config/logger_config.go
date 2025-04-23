// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/combiningwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/nonewriter"
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

	humanWarn, cfgErr := d.HumanWarnWriter.process(c, false)
	if cfgErr != nil {
		return cfgErr
	}

	humanErr, cfgErr := d.HumanErrWriter.process(c, false)
	if cfgErr != nil {
		return cfgErr
	}

	machineWarn, cfgErr := d.MachineWarnWriter.process(c, true)
	if cfgErr != nil {
		return cfgErr
	}

	machineErr, cfgErr := d.MachineErrWriter.process(c, true)
	if cfgErr != nil {
		return cfgErr
	}

	logWarn := append(humanWarn, machineWarn...)
	logErr := append(humanErr, machineErr...)

	if len(logWarn) == 0 {
		_, err := logger.SetWarnWriter(nonewriter.NewNoneWriter())
		if err != nil {
			return configerror.NewErrorf("set warn writer error: %s", err.Error())
		}
	} else if len(logWarn) == 1 {
		_, err := logger.SetWarnWriter(logWarn[0])
		if err != nil {
			return configerror.NewErrorf("set warn writer error: %s", err.Error())
		}
	} else {
		combiningWriter := combiningwriter.NewCombiningWriter(logWarn...)
		_, err := logger.SetWarnWriter(combiningWriter)
		if err != nil {
			return configerror.NewErrorf("set warn combining writer error: %s", err.Error())
		}
	}

	if len(logErr) == 0 {
		_, err := logger.SetErrWriter(nonewriter.NewNoneWriter())
		if err != nil {
			return configerror.NewErrorf("set error writer error: %s", err.Error())
		}
	} else if len(logErr) == 1 {
		_, err := logger.SetErrWriter(logErr[0])
		if err != nil {
			return configerror.NewErrorf("set error writer error: %s", err.Error())
		}
	} else {
		combiningWriter := combiningwriter.NewCombiningWriter(logErr...)
		_, err := logger.SetErrWriter(combiningWriter)
		if err != nil {
			return configerror.NewErrorf("set error combining writer error: %s", err.Error())
		}
	}

	return nil
}
