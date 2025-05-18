// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package packagelist

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var getErr error
var packageMap map[string]string

const cmdPackage = "/src/cmd"

func GetPackageList(gomod string) (map[string]string, error) {
	once.Do(func() {
		packageMap, getErr = getPackageList(gomod)
	})

	return packageMap, getErr
}

func getPackageList(gomod string) (map[string]string, error) {
	res := make(map[string]string, 5)
	cmdPackagePath := "." + cmdPackage

	count := 0
	err := filepath.Walk(cmdPackagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			templog.TempLogf("walk error: %s, skip %s\n", err.Error(), path)
			return nil
		}

		name := info.Name()

		if !info.IsDir() {
			return nil
		} else if count == 0 && name == "cmd" {
			return nil
		}

		count++

		res[name] = gomod + cmdPackage + "/" + name
		return filepath.SkipDir
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s error: %v\n", cmdPackagePath, err)
	}

	return res, nil
}
