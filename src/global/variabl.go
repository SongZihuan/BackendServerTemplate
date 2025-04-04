// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package global

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"time"
)

var Version = resource.Version
var Name = resource.Name

var Location *time.Location
