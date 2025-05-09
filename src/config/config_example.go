// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ExampleConfig struct {
	Test string `json:"test" yaml:"test" mapstructure:"test"`
}

func (d *ExampleConfig) init(filePath string, provider configparser.ConfigParserProvider) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleConfig) setDefault(c *configInfo) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleConfig) check(c *configInfo) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleConfig) process(c *configInfo) (cfgErr configerror.Error) {
	return nil
}
