// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example3

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"time"
)

type ServerExample3Core struct {
	name   string
	server *server.Server
	ctx    *servercontext.ServerContext
}

type ServerExample3CoreOption struct {
}

func NewServerExample3Core(opt *ServerExample3CoreOption) (*ServerExample3Core, error) {
	if opt == nil {
		opt = &ServerExample3CoreOption{}
	}

	sc := &ServerExample3Core{
		name: "Example-3",
	}

	return sc, nil
}

func (sc *ServerExample3Core) Name() string {
	return sc.name
}

func (sc *ServerExample3Core) Init(s *server.Server, ctx *servercontext.ServerContext) error {
	sc.server = s
	sc.ctx = ctx
	return nil
}

func (sc *ServerExample3Core) Run() error {
	if sc.server == nil || sc.ctx == nil {
		logger.Panicf("server core not init")
	}

MainCycle:
	for {
		logger.Warnf("Example3: I am running!")
		select {
		case <-sc.ctx.Listen():
			break MainCycle
		case <-time.After(2 * time.Second):
			continue
		}
	}
	return nil
}

func (sc *ServerExample3Core) Stop() {
	logger.Warnf("Example3: I am stop!")
}
