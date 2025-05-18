// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package platformlist

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/temporary/internal/templog"
	"github.com/SongZihuan/BackendServerTemplate/utils/runtimeutils"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var once sync.Once
var getErr error
var goosMap map[string]bool
var archMap map[string]bool
var platformMap map[string]map[string]bool

func GetPlatformList() (map[string]bool, map[string]bool, map[string]map[string]bool, error) {
	once.Do(func() {
		goosMap, archMap, platformMap, getErr = getPlatformList()
	})

	return goosMap, archMap, platformMap, getErr
}

func getPlatformList() (map[string]bool, map[string]bool, map[string]map[string]bool, error) {
	goos := make(map[string]bool, 5)
	goarch := make(map[string]bool, 10)
	platform := make(map[string]map[string]bool, 5)

	src := "./BUILD"

	count := 0
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk error: %s, skip %s\n", err.Error(), path)
		}

		name := info.Name()

		if info.IsDir() {
			if count == 0 && name == "BUILD" {
				return nil
			} else {
				return filepath.SkipDir
			}
		}

		count++

		if !strings.HasPrefix(name, "BUILD.") {
			templog.TempLogf("walk error: not prefix `BUILD.` , skip %s\n", path)
			return nil
		} else if !strings.HasSuffix(name, ".yaml") {
			templog.TempLogf("walk error: not suffix `.yaml` , skip %s\n", path)
			return nil
		}

		osarch := strings.Split(strings.TrimPrefix(strings.TrimSuffix(name, ".yaml"), "BUILD."), ".")

		if len(osarch) != 2 && len(osarch) != 3 {
			templog.TempLogf("walk error: unknown file name, skip %s\n", path)
			return nil
		}

		maparch, ok := runtimeutils.ServerOSArch[osarch[0]]
		if !ok {
			templog.TempLogf("walk error: bad os %s , skip %s\n", osarch[0], path)
			return nil
		}

		yes, ok := maparch[osarch[1]]
		if !yes || !ok {
			templog.TempLogf("walk error: bad arch [%s] on os [%s] , skip %s\n", osarch[1], osarch[0], path)
			return nil
		}

		if len(osarch) == 3 {
			if osarch[0] != runtimeutils.Windows {
				templog.TempLogf("walk error: bad os %s with admin mode , skip %s\n", osarch[0], path)
				return nil
			}

			if osarch[2] != "admin" {
				templog.TempLogf("walk error: unknown file name, skip %s\n", path)
				return nil
			}
		}

		if _, ok := platform[osarch[0]]; !ok {
			platform[osarch[0]] = make(map[string]bool, 5)
		}

		goos[osarch[0]] = true
		goarch[osarch[1]] = true
		platform[osarch[0]][osarch[1]] = true

		return nil
	})
	if err != nil {
		return nil, nil, nil, err
	}

	return goos, goarch, platform, nil
}
