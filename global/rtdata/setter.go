// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rtdata

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/osutils"
	"time"
)

func SetName(cfgName *string, autoName *bool, packageName string) error {
	if packageName == "" {
		return fmt.Errorf("package name is empty")
	}

	if autoName != nil && !(*autoName) { // 如果 AutoName 设置为 false, 则等同于没设置
		autoName = nil
	}

	if autoName == nil && cfgName == nil { // 采用package name
		rtdata.Name = packageName
	} else if autoName != nil && cfgName == nil { // 采用可执行文件名
		rtdata.Name = osutils.GetArgs0Name()
		if rtdata.Name == "" {
			rtdata.Name = packageName
		}
	} else if autoName == nil { // cfgName != nil 必然成立
		if *cfgName == "" {
			return fmt.Errorf("name can not be empty")
		} else if newName := cleanstringutils.GetName(*cfgName); newName != *cfgName {
			return fmt.Errorf("name is invalid: use %s please", newName)
		} else {
			rtdata.Name = *cfgName
		}
	} else { // autoName != nil && cfgName != nil 必然成立
		return fmt.Errorf("cannot specify a name and use AutoName at the same time")
	}

	if rtdata.Name == "" {
		panic("rtdata.name is empty")
	}

	return nil
}

func SetLocation(loc *time.Location) error {
	rtdata.Location = loc
	return nil
}
