// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/combiningwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/nonewriter"
	"github.com/SongZihuan/BackendServerTemplate/utils/typeutils"
)

type LoggerConfig struct {
	LogLevel loglevel.LoggerLevel `json:"log-level" yaml:"log-level" mapstructure:"log-level"`
	LogTag   typeutils.StringBool `json:"log-tag" yaml:"log-tag" mapstructure:"log-tag"`

	WarnWriter []*LoggerWriterConfig `json:"warn-writer" yaml:"warn-writer" mapstructure:"warn-writer"`
	ErrWriter  []*LoggerWriterConfig `json:"err-writer" yaml:"err-writer" mapstructure:"err-writer"`
}

func (d *LoggerConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	for _, w := range d.WarnWriter {
		err = w.init(filePath, provider)
		if err != nil && err.IsError() {
			return err
		}
	}

	for _, w := range d.ErrWriter {
		err = w.init(filePath, provider)
		if err != nil && err.IsError() {
			return err
		}
	}

	return nil
}

func (d *LoggerConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.LogLevel == "" && c.data.GlobalConfig.IsRelease() {
		d.LogLevel = loglevel.LevelInfo
	} else if d.LogLevel == "" {
		d.LogLevel = loglevel.LevelDebug
	}

	d.LogLevel = d.LogLevel.ToLower()

	if c.data.GlobalConfig.IsRelease() {
		d.LogTag.SetDefaultDisable()
	} else {
		d.LogTag.SetDefaultEnable()
	}

	for _, w := range d.WarnWriter {
		err = w.setDefault(c)
		if err != nil && err.IsError() {
			return err
		}
	}

	for _, w := range d.ErrWriter {
		err = w.setDefault(c)
		if err != nil && err.IsError() {
			return err
		}
	}

	return nil
}

func (d *LoggerConfig) check(c *configInfo) (err configerror.Error) {
	if !d.LogLevel.OK() {
		return configerror.NewErrorf("log level error: %s", d.LogLevel)
	} else if d.LogLevel == loglevel.PseudoLevelTag {
		return configerror.NewErrorf("log level error: %s", loglevel.PseudoLevelTag)
	}

	for _, w := range d.WarnWriter {
		err = w.check(c)
		if err != nil && err.IsError() {
			return err
		}
	}

	for _, w := range d.ErrWriter {
		err = w.check(c)
		if err != nil && err.IsError() {
			return err
		}
	}

	return nil
}

func (d *LoggerConfig) process(c *configInfo) configerror.Error {
	logWarn := make([]logwriter.Writer, 0, len(d.WarnWriter))
	logErr := make([]logwriter.Writer, 0, len(d.ErrWriter))

	for _, w := range d.WarnWriter {
		writer, cfgErr := w.process(c)
		if cfgErr != nil && cfgErr.IsError() {
			return cfgErr
		}

		logWarn = append(logWarn, writer)
	}

	for _, w := range d.ErrWriter {
		writer, cfgErr := w.process(c)
		if cfgErr != nil && cfgErr.IsError() {
			return cfgErr
		}

		logErr = append(logErr, writer)
	}

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
