// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/server/serverinterface"
	"sync"
	"time"
)

type ServerExample1 struct {
	running      bool
	ctx          *servercontext.ServerContext
	name         string
	wg           *sync.WaitGroup
	stopWaitTime time.Duration
}

type ServerExample1Option struct {
	StopWaitTime time.Duration
}

func NewServerExample1(opt *ServerExample1Option) (*ServerExample1, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ServerExample1Option{
			StopWaitTime: 10 * time.Second,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	server := &ServerExample1{
		ctx:          ctx,
		running:      false,
		name:         "example1",
		wg:           new(sync.WaitGroup),
		stopWaitTime: opt.StopWaitTime,
	}
	err := server.init()
	if err != nil {
		return nil, nil, err
	}

	return server, ctx, nil
}

func (s *ServerExample1) init() error {
	return nil
}

func (s *ServerExample1) Name() string {
	return s.name
}

func (s *ServerExample1) GetCtx() *servercontext.ServerContext {
	return s.ctx
}

func (s *ServerExample1) Run() {
	s.running = true
	defer func() {
		s.running = false
	}()

	s.wg = new(sync.WaitGroup)
	s.wg.Add(1)
	defer s.wg.Done()

MainCycle:
	for {
		fmt.Println("Example1: I am running!")

		select {
		case <-s.ctx.Listen():
			fmt.Println("Example1: I am stop!")
			break MainCycle
		case <-time.After(1 * time.Second):
			continue
		}
	}
}

func (s *ServerExample1) Stop() {
	s.ctx.StopTask()
	if s.wg != nil {
		wgchan := make(chan any)

		go func() {
			s.wg.Wait()
			close(wgchan)
		}()

		select {
		case <-time.After(s.stopWaitTime):
		case <-wgchan:
		}
	}
}

func (s *ServerExample1) IsRunning() bool {
	return s.running
}

func _test() {
	var a serverinterface.Server
	var b *ServerExample1

	a = b
	_ = a
}
