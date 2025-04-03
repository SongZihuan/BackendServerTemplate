// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"encoding/json"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"os"
	"reflect"
)

type JsonProvider struct {
	HasRead  bool
	FileData []byte
}

func NewJsonProvider() *JsonProvider {
	return &JsonProvider{
		HasRead:  false,
		FileData: nil,
	}
}

func (j *JsonProvider) CanUTF8() bool {
	return true
}

func (j *JsonProvider) ReadFile(filepath string) configerror.Error {
	if j.HasRead {
		return configerror.NewErrorf("config file has been read")
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return configerror.NewErrorf(fmt.Sprintf("read file error: %s", err.Error()))
	}

	j.FileData = data
	j.HasRead = true

	return nil
}

func (j *JsonProvider) ParserFile(target any) configerror.Error {
	if !j.HasRead || j.FileData == nil {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	err := json.Unmarshal(j.FileData, target)
	if err != nil {
		return configerror.NewErrorf("json parser error: %s", err.Error())
	}

	return nil
}

func (j *JsonProvider) WriteFile(filepath string, src any) configerror.Error {
	if !j.HasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	target, err := json.Marshal(src)
	if err != nil {
		return configerror.NewErrorf("json marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, target, 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}

func _testJson() {
	var a ConfigParserProvider
	a = &JsonProvider{}
	_ = a
}
