// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/utils/strconvutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/typeutils"
	"time"
)

type ServerExampleConfig struct {
	StopWaitTime    string               `json:"stop-wait-time" yaml:"stop-wait-time" mapstructure:"stop-wait-time"`
	StartupWaitTime string               `json:"startup-wait-time" yaml:"startup-wait-time" mapstructure:"startup-wait-time"`
	LockThread      typeutils.StringBool `json:"lock-thread" yaml:"lock-thread" mapstructure:"lock-thread"`

	StopWaitTimeDuration    time.Duration `json:"-" yaml:"-" mapstructure:"-"`
	StartupWaitTimeDuration time.Duration `json:"-" yaml:"-" mapstructure:"-"`
}

func (d *ServerExampleConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *ServerExampleConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.StartupWaitTime == "" {
		d.StartupWaitTime = "3s"
	}

	if d.StopWaitTime == "" {
		d.StopWaitTime = "10s"
	}

	d.LockThread.SetDefaultDisable()
	return nil
}

func (d *ServerExampleConfig) check(c *configInfo) (err configerror.Error) {
	return nil
}

func (d *ServerExampleConfig) process(c *configInfo) (cfgErr configerror.Error) {
	d.StopWaitTimeDuration = strconvutils.ReadTimeDurationPositive(d.StopWaitTime) // 不可能小于0
	if d.StopWaitTimeDuration < 10*time.Second {
		configerror.ShowWarningf("stop-wait-time (value: %s) is less than the recommended value of 10s", strconvutils.TimeDurationToString(d.StopWaitTimeDuration))
	}

	d.StartupWaitTimeDuration = strconvutils.ReadTimeDurationPositive(d.StartupWaitTime) // 不可能小于0
	if d.StartupWaitTimeDuration < 3*time.Second {
		configerror.ShowWarningf("startup-wait-time (value: %s) is less than the recommended value of 3s", strconvutils.TimeDurationToString(d.StopWaitTimeDuration))
	}

	return nil
}
