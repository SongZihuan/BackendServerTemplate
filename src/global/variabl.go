// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package global

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"time"
)

var (
	// Version SemanticVersioning License Report 继承自resource
	Version            = resource.Version
	SemanticVersioning = resource.SemanticVersioning
	License            = resource.License
	Report             = resource.Report
	BuildTime          = resource.BuildTime
	GitCommitHash      = resource.GitCommitHash
	GitTag             = resource.GitTag
	GitTagCommitHash   = resource.GitTagCommitHash

	// Name 当config读取配置文件加载时可能会被更改
	Name = resource.Name
)

var (
	Location *time.Location
)
