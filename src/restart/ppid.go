// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package restart

import (
	"github.com/SongZihuan/BackendServerTemplate/utils/processutils"
	"time"
)

func PpidWatcher(ppid int) chan any {
	ppidchan := make(chan any)

	if ppid == 0 {
		return ppidchan
	}

	go func() {
	PpidCycle:
		for range time.Tick(3 * time.Second) {
			if !processutils.PidExists(ppid) {
				close(ppidchan)
				break PpidCycle
			}
		}
	}()

	return ppidchan
}
