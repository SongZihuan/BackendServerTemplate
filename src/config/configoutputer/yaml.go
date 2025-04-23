// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configoutputer

import (
	"bytes"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

type YamlProvider struct {
	Ident int
}

func NewYamlProvider(opt *NewConfigOutputProviderOption) *YamlProvider {
	if opt == nil {
		opt = new(NewConfigOutputProviderOption)
	}

	if opt.Ident <= 0 {
		opt.Ident = 4
	}

	return &YamlProvider{
		Ident: opt.Ident,
	}
}

func (y *YamlProvider) CanUTF8() bool {
	return true
}

func (y *YamlProvider) WriteFile(filepath string, src any) configerror.Error {
	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	var buf bytes.Buffer

	encoder := yaml.NewEncoder(&buf)
	defer func() {
		_ = encoder.Close()
	}()

	encoder.SetIndent(4) // 设置缩进长度

	err := encoder.Encode(src)
	if err != nil {
		return configerror.NewErrorf("yaml marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, buf.Bytes(), 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}

func _testYaml() {
	var a ConfigOutputProvider
	a = &YamlProvider{}
	_ = a
}
