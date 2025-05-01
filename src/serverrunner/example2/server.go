// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example2

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"time"
)

type ServerExample2Core struct {
	name   string
	server *server.Server
	ctx    *servercontext.ServerContext
}

type ServerExample2CoreOption struct {
}

func NewServerExample2Core(opt *ServerExample2CoreOption) (*ServerExample2Core, error) {
	if opt == nil {
		opt = &ServerExample2CoreOption{}
	}

	sc := &ServerExample2Core{
		name: "Example-2",
	}

	return sc, nil
}

func (sc *ServerExample2Core) Name() string {
	return sc.name
}

func (sc *ServerExample2Core) Init(s *server.Server, ctx *servercontext.ServerContext) error {
	sc.server = s
	sc.ctx = ctx
	return nil
}

func (sc *ServerExample2Core) Run() error {
	if sc.server == nil || sc.ctx == nil {
		logger.Panicf("server core not init")
	}

MainCycle:
	for {
		fmt.Printf("Example2: Hello, %s. I am running!\n", config.Data().Server.Name)

		select {
		case <-sc.ctx.Listen():
			break MainCycle
		case <-time.After(2 * time.Second):
			continue
		}
	}

	return nil
}

func (sc *ServerExample2Core) Stop() {
	fmt.Println("Example2: I am stop!")
}
