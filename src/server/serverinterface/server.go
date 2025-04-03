// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package serverinterface

import (
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
)

const ControllerName = "controller"

type Server interface {
	Name() string
	Run()
	GetCtx() *servercontext.ServerContext
	Stop()
	IsRunning() bool
}
