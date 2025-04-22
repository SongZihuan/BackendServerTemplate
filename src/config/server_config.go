// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"time"
)

type ServerConfig struct {
	StopWaitTime string `json:"stop-wait-time" yaml:"stop-wait-time" mapstructure:"stop-wait-time"`
	Name         string `json:"name" yaml:"name" mapstructure:"name"`

	StopWaitTimeDuration time.Duration `json:"-" yaml:"-" mapstructure:"-"`
}

func (d *ServerConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *ServerConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.StopWaitTime == "" {
		d.StopWaitTime = "10s"
	}

	if d.Name == "" {
		d.Name = "Jack"
	}
	return nil
}

func (d *ServerConfig) check(c *configInfo) (err configerror.Error) {
	return nil
}

func (d *ServerConfig) process(c *configInfo) (cfgErr configerror.Error) {
	d.StopWaitTimeDuration = strconvutils.ReadTimeDurationPositive(d.StopWaitTime)
	if d.StopWaitTimeDuration < 10*time.Second {
		_ = configerror.NewWarningf("stop-wait-time (value: %s) is less than the recommended value of 10s", strconvutils.TimeDurationToString(d.StopWaitTimeDuration))
	}

	return nil
}
