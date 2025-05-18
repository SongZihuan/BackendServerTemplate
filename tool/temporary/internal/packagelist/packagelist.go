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
var packageList []string

func GetPackageList() ([]string, error) {
	once.Do(func() {
		packageList, getErr = getPackageList()
	})

	return packageList, getErr
}

func getPackageList() ([]string, error) {
	res := make([]string, 0, 5)
	src := "./src/cmd"

	count := 0
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
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

		res = append(res, name)
		return filepath.SkipDir
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s error: %v\n", src, err)
	}

	return res, nil
}
