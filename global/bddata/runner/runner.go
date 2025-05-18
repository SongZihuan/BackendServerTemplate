// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runner

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/bdmodule"
	"time"
)

var data bdmodule.GlobalData
var config *bdmodule.BuildConfigData

const defaultPackageName = "default"

func ReadGlobalData(dat []byte, packageName string) (err error) {
	decoder := gob.NewDecoder(bytes.NewReader(dat))
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("error decoding data: %s", err.Error())
	}

	cfg, ok := data.ConfigSet[packageName]
	if !ok || cfg == nil {
		cfg, ok = data.ConfigSet[defaultPackageName]
		if !ok || cfg == nil {
			cfg = new(bdmodule.BuildConfigData)
		}
	}

	config = cfg

	return nil
}

func GetLongVersion() string {
	return data.LongVersion
}

func GetLongSemanticVersion() string {
	return data.LongSemanticVersion
}

func GetShortVersion() string {
	return data.ShortVersion
}

func GetShortSemanticVersion() string {
	return data.ShortSemanticVersion
}

func GetLicense() string {
	return data.License
}

func GetReport() string {
	return data.Report
}

func GetModuleName() string {
	return data.ModuleMame
}

func GetCommitHash() string {
	return data.CommitHash
}

func GetBuildDate() time.Time {
	return data.BuildDate
}

func GetConfigName() string {
	if config.Name == nil {
		return ""
	}
	return *(config.Name)
}

func GetConfigNamePointer() *string {
	if config.Name == nil {
		return nil
	}

	var tmp = *(config.Name)
	return &tmp
}

func GetConfigAutoName() bool {
	if config.AutoName == nil {
		return false
	}
	return *(config.AutoName)
}

func GetConfigAutoNamePointer() *bool {
	if config.AutoName == nil {
		return nil
	}

	var tmp = *(config.AutoName)
	return &tmp
}

func GetConfigEnvPrefix() string {
	return config.EnvPrefix
}

func GetConfigService() (*bdmodule.ServiceConfig, error) {
	return config.Service.Copy()
}

func GetConfigServiceMust() *bdmodule.ServiceConfig {
	res, err := config.Service.Copy()
	if err != nil {
		return nil
	}
	return res
}
