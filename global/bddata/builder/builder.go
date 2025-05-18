// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package builder

import (
	"encoding/gob"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/bdmodule"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var data bdmodule.GlobalData

func SaveGlobalData(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&data) // 传入指针，但实际编码的数据是结构体
	if err != nil {
		return fmt.Errorf("error encoding data: %s", err.Error())
	}

	return nil
}

func SetLongVersion(longVersion, longSemanticVersion string) {
	data.LongVersion = longVersion
	data.LongSemanticVersion = longSemanticVersion
}

func SetShortVersion(shortVersion, shortSemanticVersion string) {
	data.ShortVersion = shortVersion
	data.ShortSemanticVersion = shortSemanticVersion
}

func SetBasicInfo(license, report string) {
	data.License = license
	data.Report = report
}

func SetModuleName(moduleMame string) {
	data.ModuleMame = moduleMame
}

func SetCommitHash(hash string) {
	data.CommitHash = hash
}

func SetBuildDate(date time.Time) {
	data.BuildDate = date
}

func SetConfig(dat []byte) error {
	var dest bdmodule.BuildConfigSet
	err := yaml.Unmarshal(dat, &dest)
	if err != nil {
		return err
	}

	err = dest.CheckAndSetDefault()
	if err != nil {
		return err
	}

	data.ConfigSet = dest

	return nil
}
