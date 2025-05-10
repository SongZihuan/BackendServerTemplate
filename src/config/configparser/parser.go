// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import "github.com/SongZihuan/BackendServerTemplate/src/config/configerror"

type ConfigParserProvider interface {
	CanUTF8() bool // Must return true
	ReadFile(filepath string) configerror.Error
	ParserFile(target any) configerror.Error
}

type NewConfigParserProviderOption struct {
	EnvPrefix  string
	AutoReload bool
}

type NewProvider func(opt *NewConfigParserProviderOption) ConfigParserProvider
