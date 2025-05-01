// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ServerConfig struct {
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	Example1   ServerExampleConfig    `json:"example1" yaml:"example1" mapstructure:"example1"`
	Example2   ServerExampleConfig    `json:"example2" yaml:"example2" mapstructure:"example2"`
	Example3   ServerExampleConfig    `json:"example3" yaml:"example3" mapstructure:"example3"`
	Controller ServerControllerConfig `json:"controller" yaml:"controller" mapstructure:"controller"`
}

func (d *ServerConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	cfgErr := d.Example1.init(filePath, provider)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example2.init(filePath, provider)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example3.init(filePath, provider)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Controller.init(filePath, provider)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	return nil
}

func (d *ServerConfig) setDefault(c *configInfo) configerror.Error {
	if d.Name == "" {
		d.Name = "Jack"
	}

	cfgErr := d.Example1.setDefault(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example2.setDefault(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example3.setDefault(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Controller.setDefault(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	return nil
}

func (d *ServerConfig) check(c *configInfo) configerror.Error {
	cfgErr := d.Example1.check(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example2.check(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example3.check(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Controller.check(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	return nil
}

func (d *ServerConfig) process(c *configInfo) configerror.Error {
	cfgErr := d.Example1.process(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example2.process(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Example3.process(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	cfgErr = d.Controller.process(c)
	if cfgErr != nil && cfgErr.IsError() {
		return cfgErr
	}

	return nil
}
