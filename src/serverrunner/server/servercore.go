// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"time"
)

const ControllerName = "controller"

type ServerCore interface {
	Name() string
	Init(server *Server, ctx *servercontext.ServerContext) error
	Run() error
	Stop()
}

type ControllerServerCore interface {
	ServerCore
	AddServerCore(option *ServerOption) (*Server, error)
	AddServer(ser *Server) error
	DelServer(ser *Server) error
	GetStopWaitTime() time.Duration
	GetStartupWaitTime() time.Duration

	Controller() // 该函数用于表示该Core属于控制器，无实际运行作用
}
