// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configoutputer

import "github.com/SongZihuan/BackendServerTemplate/src/config/configerror"

type ConfigOutputProvider interface {
	CanUTF8() bool // Must return true
	WriteFile(filepath string, data any) configerror.Error
}

type NewConfigOutputProviderOption struct {
	Ident int
}

type NewProvider func(opt *NewConfigOutputProviderOption) ConfigOutputProvider
