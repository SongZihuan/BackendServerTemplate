// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configoutputer

import (
	"encoding/json"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"os"
	"reflect"
	"strings"
)

type JsonProvider struct {
	Ident string
}

func NewJsonProvider(opt *NewConfigOutputProviderOption) *JsonProvider {
	if opt == nil {
		opt = new(NewConfigOutputProviderOption)
	}

	if opt.Ident <= 0 {
		opt.Ident = 4
	}

	return &JsonProvider{
		Ident: strings.Repeat(" ", opt.Ident),
	}
}

func (j *JsonProvider) CanUTF8() bool {
	return true
}

func (j *JsonProvider) WriteFile(filepath string, src any) configerror.Error {
	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	target, err := json.MarshalIndent(src, "", j.Ident)
	if err != nil {
		return configerror.NewErrorf("json marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, target, 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}
