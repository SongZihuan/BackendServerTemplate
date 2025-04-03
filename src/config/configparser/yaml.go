// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

type YamlProvider struct {
	HasRead  bool
	FileData []byte
}

func NewYamlProvider() *YamlProvider {
	return &YamlProvider{
		HasRead:  false,
		FileData: nil,
	}
}

func (y *YamlProvider) CanUTF8() bool {
	return true
}

func (y *YamlProvider) ReadFile(filepath string) configerror.Error {
	if y.HasRead {
		return configerror.NewErrorf("config file has been read")
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return configerror.NewErrorf(fmt.Sprintf("read file error: %s", err.Error()))
	}

	y.FileData = data
	y.HasRead = true

	return nil
}

func (y *YamlProvider) ParserFile(target any) configerror.Error {
	if !y.HasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	err := yaml.Unmarshal(y.FileData, target)
	if err != nil {
		return configerror.NewErrorf("yaml unmarshal error: %s", err.Error())
	}

	return nil
}

func (y *YamlProvider) WriteFile(filepath string, src any) configerror.Error {
	if !y.HasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	target, err := yaml.Marshal(src)
	if err != nil {
		return configerror.NewErrorf("yaml marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, target, 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}

func _testYaml() {
	var a ConfigParserProvider
	a = &YamlProvider{}
	_ = a
}
