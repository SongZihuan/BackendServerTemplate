// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mod

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"os"
	"regexp"
	"strings"
	"sync"
)

const FileGoMod = "./go.mod"
const module = "module"

var validModuleNameRegex = regexp.MustCompile(`^[a-zA-Z0-9]+([./-]?[a-zA-Z0-9]+)*$`)

var once sync.Once
var goModuleName string = ""
var onceErr error

func InitGoModuleName() error {
	_, err := GetGoModuleName()
	return err
}

func GetGoModuleName() (string, error) {
	once.Do(func() {
		goModuleName, onceErr = getGoModuleName()
	})

	return goModuleName, onceErr
}

func getGoModuleName() (string, error) {
	genlog.GenLog("find the go mod name")
	defer genlog.GenLog("find the go mod name finish")

	dat, err := os.ReadFile(FileGoMod)
	if err != nil {
		return "", err
	}

	goMod := strings.TrimPrefix(cleanstringutils.GetString(string(dat)), "\n")

	moduleLine := strings.TrimSpace(strings.Split(goMod, "\n")[0])

	if !strings.Contains(moduleLine, module) && len(moduleLine) > len(module) {
		return "", fmt.Errorf("go.mod error: %s not found", module)
	}

	moduleName := cleanstringutils.GetStringOneLine(moduleLine[len(module):])
	if !isValidGoModuleName(moduleName) {
		return "", fmt.Errorf("go.mod error: '%s' is not a valid go module name", moduleName)
	}

	genlog.GenLogf("go mod name get: %s\n", moduleName)
	return moduleName, nil
}

func isValidGoModuleName(name string) bool {
	return validModuleNameRegex.MatchString(name)
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
