// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"sync"
	"time"
)

func Run(r Runner) (err error, timeout bool) {
	errchan := make(chan error, 1)
	startupTime := r.StartupWaitTime()
	go r.Run(errchan)

	select {
	case err, ok := <-errchan:
		if ok && err != nil {
			return err, false // 视为启动失败
		}
		return nil, false
	case <-time.After(startupTime):
		return nil, true
	}
}

func RunWithWorkGroup(r Runner, wg *sync.WaitGroup) (err error, timeout bool) {
	errchan := make(chan error, 1)
	startupTime := r.StartupWaitTime()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Run(errchan)
	}()

	select {
	case err, ok := <-errchan:
		if ok && err != nil {
			return err, false // 视为启动失败
		}
		return nil, false
	case <-time.After(startupTime):
		return nil, true
	}
}
