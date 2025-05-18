// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package modutils

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"os"
	"regexp"
	"strings"
)

const FileGoMod = "./go.mod"
const module = "module"

var validModuleNameRegex = regexp.MustCompile(`^[a-zA-Z0-9]+([./-]?[a-zA-Z0-9]+)*$`)

func GetGoModuleName() (string, error) {
	dat, err := os.ReadFile(FileGoMod)
	if err != nil {
		return "", err
	}

	goMod := strings.TrimPrefix(cleanstringutils.GetString(string(dat)), "\n")

	moduleLine := strings.TrimSpace(strings.Split(goMod, "\n")[0])

	if !strings.Contains(moduleLine, module) && len(moduleLine) > len(module) {
		return "", fmt.Errorf("go.mod error: %s not found", module)
	}

	moduleName := strings.Trim(cleanstringutils.GetStringOneLine(moduleLine[len(module):]), "/")
	if !isValidGoModuleName(moduleName) {
		return "", fmt.Errorf("go.mod error: '%s' is not a valid go module name", moduleName)
	}

	return moduleName, nil
}

func isValidGoModuleName(name string) bool {
	return validModuleNameRegex.MatchString(name)
}
