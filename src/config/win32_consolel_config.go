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

type Win32ConsoleConfig struct {
	UseOn                string               `json:"use-on" yaml:"use-on" mapstructure:"use-on"`
	Use                  bool                 `json:"-" yaml:"-" mapstructure:"-"`
	CtrlCExit            typeutils.StringBool `json:"ctrl-c-exit" yaml:"ctrl-c-exit" mapstructure:"ctrl-c-exit"`
	CtrlBreakExit        typeutils.StringBool `json:"ctrl-break-exit" yaml:"ctrl-break-exit" mapstructure:"ctrl-break-exit"`
	ConsoleCloseRecovery typeutils.StringBool `json:"console-close-recovery" yaml:"console-close-recovery" mapstructure:"console-close-recovery"`
}

func (d *Win32ConsoleConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *Win32ConsoleConfig) setDefault(c *configInfo) (err configerror.Error) {
	if d.UseOn == "" {
		d.UseOn = "only-win32"
	}

	d.CtrlCExit.SetDefaultEnable()
	d.CtrlBreakExit.SetDefaultEnable()
	d.ConsoleCloseRecovery.SetDefaultDisable()

	return nil
}

func (d *Win32ConsoleConfig) check(c *configInfo) (err configerror.Error) {
	if d.UseOn != "any" && d.UseOn != "not-win32" && d.UseOn != "only-win32" && d.UseOn != "never" {
		return configerror.NewErrorf("bad use-on: %s, must be one of (any, not-win32, only-win32, never)", d.UseOn)
	}
	return nil
}

func (d *Win32ConsoleConfig) process(c *configInfo) (cfgErr configerror.Error) {
	switch d.UseOn {
	case "any", "only-win32":
		d.Use = runtime.GOOS == "windows"
	case "never", "not-win32":
		d.Use = false
	default:
		logger.Panic("error use-on!") // 正常情况下，非正确值应该在check步骤被返回，若此处发现错误值则可能是check的逻辑有误
	}
	return nil
}
