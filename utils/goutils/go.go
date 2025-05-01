// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package goutils

import (
	"fmt"
	"runtime"
)

var DefaultMaxProcs = runtime.GOMAXPROCS(0)

func LockOSThread() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("lock thread failed: %v", err)
		}
	}()

	runtime.GOMAXPROCS(runtime.GOMAXPROCS(0) + 1)
	runtime.LockOSThread()
	return nil
}

func UnlockOSThread() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("unlock thread failed: %v", err)
		}
	}()

	procs := runtime.GOMAXPROCS(0) - 1
	if procs < DefaultMaxProcs {
		procs = DefaultMaxProcs
	}

	runtime.GOMAXPROCS(procs)
	runtime.UnlockOSThread()
	return nil
}
