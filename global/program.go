// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package global

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/buildinfo"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/runner"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
)

func ProgramInit(packageName string) error {
	if packageName == "" {
		return fmt.Errorf("package name is empty")
	} else if n := cleanstringutils.GetString(packageName); packageName != n {
		return fmt.Errorf("package name is invalid, use %s please", n)
	}

	err := runner.ReadGlobalData(buildinfo.BuildData, packageName)
	if err != nil {
		return err
	}

	err = rtdata.SetName(runner.GetConfig().Name, runner.GetConfig().AutoName, packageName)
	if err != nil {
		return err
	}

	return nil
}

func get[T any](val *T) T {
	if val == nil {
		var t T
		return t
	}
	return *val
}
