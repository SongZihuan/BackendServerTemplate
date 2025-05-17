// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

type ControllerContext interface {
	StartupRun() // 进入startup后开始运行
	StartRun()   // go协程进行时首先运行
	FinishRun()  // go协程推出后defer运行
}
