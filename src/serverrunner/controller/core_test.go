// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"testing"
)

func TestControllerCore(t *testing.T) {
	var a server.ControllerServerCore
	var b server.ServerCore
	var c *ControllerCore

	a = c
	b = c
	_ = a
	_ = b
}
