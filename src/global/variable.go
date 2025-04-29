// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package global

import (
	"fmt"
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/timeutils"
	"strings"
	"time"
)

var (
	// Version SemanticVersioning License Report BuildTime GitCommitHash GitTag GitTagCommitHash EnvPrefix 继承自resource（程序init完成后即可调用）
	Version            = resource.Version
	SemanticVersioning = resource.SemanticVersioning
	License            = resource.License
	Report             = resource.Report
	BuildTime          = resource.BuildTime
	GitCommitHash      = resource.GitCommitHash
	GitTag             = resource.GitTag
	GitTagCommitHash   = resource.GitTagCommitHash
	EnvPrefix          = resource.EnvPrefix

	// Name 继承自resource
	// 注意：命令行参数或配置文件加载时可能会被更改
	Name            = resource.Name
	NameFlagChanged = false
)

// Location 以下变量需要在配置文件加载完毕后才可调用
var (
	UTCLocation   = time.UTC
	LocalLocation = timeutils.GetLocalTimezone()
	Location      = time.UTC
)

func init() {
	if EnvPrefix == "" {
		return
	}

	newEnvPrefix := envutils.ToEnvName(EnvPrefix)
	if EnvPrefix != newEnvPrefix {
		panic(fmt.Errorf("bad %s; good %s", EnvPrefix, newEnvPrefix))
	} else if strings.HasSuffix(EnvPrefix, "_") {
		panic("EnvPrefix End With '_'")
	}
}
