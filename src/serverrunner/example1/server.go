// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"time"
)

type ServerExample1Core struct {
	name   string
	server *server.Server
	ctx    *servercontext.ServerContext
}

type ServerExample1CoreOption struct {
}

func NewServerExample1Core(opt *ServerExample1CoreOption) (*ServerExample1Core, error) {
	if opt == nil {
		opt = &ServerExample1CoreOption{}
	}

	sc := &ServerExample1Core{
		name: "Example-1",
	}

	return sc, nil
}

func (sc *ServerExample1Core) Name() string {
	return sc.name
}

func (sc *ServerExample1Core) Init(s *server.Server, ctx *servercontext.ServerContext) error {
	sc.server = s
	sc.ctx = ctx
	return nil
}

func (sc *ServerExample1Core) Run() error {
	if sc.server == nil || sc.ctx == nil {
		logger.Panicf("server core not init")
	}

MainCycle:
	for {
		if global.GitTag == "" || global.GitTagCommitHash == "" {
			fmt.Printf("Example1: I am running! BuildDate: '%s' Commit: '%s' Version: '%s' Now: '%s'\n", global.BuildTime.Format(time.DateTime), global.GitCommitHash, global.Version, time.Now().Format(time.DateTime))
		} else {
			fmt.Printf("Example1: I am running! BuildDate: '%s' Commit: '%s' Tag: '%s' Tag Commit: '%s' Version: '%s' Now: '%s'\n", global.BuildTime.Format(time.DateTime), global.GitCommitHash, global.GitTag, global.GitTagCommitHash, global.Version, time.Now().Format(time.DateTime))
		}

		select {
		case <-sc.ctx.Listen():
			break MainCycle
		case <-time.After(1 * time.Second):
			continue
		}
	}

	return nil
}

func (sc *ServerExample1Core) Stop() {
	fmt.Println("Example1: I am stop!")
}
