// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"time"
)

func Run(r Runner) (err error, timeout bool) {
	startupErr := make(chan RunStartupError, 1)
	startupTime := r.StartupWaitTime()

	go r.Run(startupErr)

	select {
	case err, ok := <-startupErr:
		if ok && err != nil {
			cleanTime := err.CleanTime()

			select {
			case <-err.Wait():
				// pass
			case <-time.After(cleanTime):
				// pass
			}

			return err.Error(), false // 视为启动失败
		}
		return nil, false // 启动成功（因为通道关闭但没传递错误）
	case <-time.After(startupTime):
		return nil, true
	}
}

func RunByController(r Runner, cctx ControllerContext) (err error, timeout bool) {
	cctx.StartupRun()

	startupErr := make(chan RunStartupError, 1)
	startupTime := r.StartupWaitTime()

	go func() {
		cctx.StartRun()
		defer cctx.FinishRun()

		r.Run(startupErr)
	}()

	select {
	case err, ok := <-startupErr:
		if ok && err != nil {
			cleanTime := err.CleanTime()

			select {
			case <-err.Wait():
				// pass
			case <-time.After(cleanTime):
				// pass
			}

			return err.Error(), false // 视为启动失败
		}
		return nil, false // 启动成功（因为通道关闭但没传递错误）
	case <-time.After(startupTime):
		return nil, true
	}
}
