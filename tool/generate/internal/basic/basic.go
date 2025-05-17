// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basic

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
)

func init() {
	resource.InitVar()
}

func WriteBasicData() error {
	genlog.GenLog("write basic data")
	defer genlog.GenLog("write basic data finish")

	builder.SetBasicInfo(resource.License, resource.Report)

	return nil
}
