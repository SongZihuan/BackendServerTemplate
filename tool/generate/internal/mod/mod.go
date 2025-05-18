// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mod

import (
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/utils/modutils"
	"sync"
)

var once sync.Once
var goModuleName string = ""
var onceErr error

func InitGoModuleName() error {
	_, err := GetGoModuleName()
	return err
}

func GetGoModuleName() (string, error) {
	once.Do(func() {
		genlog.GenLog("find the go mod name")
		defer genlog.GenLog("find the go mod name finish")

		goModuleName, onceErr = modutils.GetGoModuleName()
		if onceErr == nil {
			genlog.GenLogf("go mod name get: %s\n", goModuleName)
		}
	})

	return goModuleName, onceErr
}

func WriteModuleNameData() (err error) {
	genlog.GenLog("write go module name data")
	defer genlog.GenLog("write go module nam finish")

	moduleMame, err := GetGoModuleName()
	if err != nil {
		genlog.GenLog("get go module info failed")
		return err
	}

	builder.SetModuleName(moduleMame)

	return nil
}
