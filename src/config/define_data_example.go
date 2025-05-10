// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ExampleData struct {
}

func (d *ExampleData) init(filePath string, provider configparser.ConfigParserProvider) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleData) setDefault(c *configInfo) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleData) check(c *configInfo) (cfgErr configerror.Error) {
	return nil
}

func (d *ExampleData) process(c *configInfo) (cfgErr configerror.Error) {
	return nil
}
