// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package processutils

import (
	"github.com/shirou/gopsutil/v4/process"
)

func PidExists(pid int) bool {
	ret, err := process.PidExists(int32(pid))
	if err != nil {
		return false
	}
	return ret
}
