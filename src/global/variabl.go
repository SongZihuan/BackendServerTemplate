// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package global

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"time"
)

var (
	// Version SemanticVersioning License Report BuildTime GitCommitHash GitTag GitTagCommitHash 继承自resource（程序init完成后即可调用）
	Version            = resource.Version
	SemanticVersioning = resource.SemanticVersioning
	License            = resource.License
	Report             = resource.Report
	BuildTime          = resource.BuildTime
	GitCommitHash      = resource.GitCommitHash
	GitTag             = resource.GitTag
	GitTagCommitHash   = resource.GitTagCommitHash

	// Name 继承自resource（程序init完成后即可调用）
	// 注意：命令行参数或配置文件加载时可能会被更改
	Name = resource.Name
)

// Location 以下变量需要在配置文件加载完毕后才可调用
var (
	Location = time.UTC
)
