// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/utils/typeutils"
	"runtime"
)

type SignalConfig struct {
	UseOn       string               `json:"use-on" yaml:"use-on" mapstructure:"use-on"`
	Use         bool                 `json:"-" yaml:"-" mapstructure:"-"`
	SigIntExit  typeutils.StringBool `json:"sigint-exit" yaml:"sigint-exit" mapstructure:"sigint-exit"`
	SigTermExit typeutils.StringBool `json:"sigterm-exit" yaml:"sigterm-exit" mapstructure:"sigterm-exit"`
	SigHupExit  typeutils.StringBool `json:"sighup-exit" yaml:"sighup-exit" mapstructure:"sighup-exit"`
	SigQuitExit typeutils.StringBool `json:"sigquit-exit" yaml:"sigquit-exit" mapstructure:"sigquit-exit"`
}

func (d *SignalConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *SignalConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.UseOn == "" {
		d.UseOn = "not-win32"
	}

	d.SigIntExit.SetDefaultEnable()
	d.SigTermExit.SetDefaultEnable()
	d.SigHupExit.SetDefaultEnable()

	if c.data.IsRelease() {
		d.SigQuitExit.SetDefaultDisable()
	} else {
		d.SigQuitExit.SetDefaultEnable()
	}

	return nil
}

func (d *SignalConfig) check(c *configInfo) (err configerror.Error) {
	if d.UseOn != "any" && d.UseOn != "not-win32" && d.UseOn != "only-win32" && d.UseOn != "never" {
		return configerror.NewErrorf("bad use-on: %s, must be one of (any, not-win32, only-win32, never)", d.UseOn)
	}
	return nil
}

func (d *SignalConfig) process(c *configInfo) (cfgErr configerror.Error) {
	switch d.UseOn {
	case "any":
		d.Use = true
	case "never":
		d.Use = false
	case "not-win32":
		d.Use = runtime.GOOS != "windows"
	case "only-win32":
		d.Use = runtime.GOOS == "windows"
	default:
		logger.Panic("error use-on!") // 正常情况下，非正确值应该在check步骤被返回，若此处发现错误值则可能是check的逻辑有误
	}
	return nil
}
