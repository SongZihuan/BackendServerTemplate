// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rtdata

import (
	"github.com/SongZihuan/BackendServerTemplate/utils/timeutils"
	"time"
)

type GlobalRuntimeData struct {
	Name string // 继承自 `GlobalData`

	UTC           *time.Location
	LocalLocation *time.Location
	Location      *time.Location
}

var rtdata GlobalRuntimeData

func init() {
	rtdata.UTC = time.UTC
	rtdata.LocalLocation = timeutils.GetLocalTimezone()
	rtdata.Location = time.UTC
}
