// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import "time"

type Runner interface {
	Run(chan error)
	StartupWaitTime() time.Duration
	Runner() // 用于标记该对象为一个Runner，无实际作用和调用
}
