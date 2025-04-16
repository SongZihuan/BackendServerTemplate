// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example3

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/server/serverinterface"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"sync"
	"time"
)

type ServerExample3 struct {
	running      bool
	ctx          *servercontext.ServerContext
	name         string
	wg           *sync.WaitGroup
	stopWaitTime time.Duration
}

type ServerExample3Option struct {
	StopWaitTime time.Duration
}

func NewServerExample3(opt *ServerExample3Option) (*ServerExample3, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ServerExample3Option{
			StopWaitTime: 10 * time.Second,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	server := &ServerExample3{
		ctx:          ctx,
		running:      false,
		name:         "example3",
		wg:           new(sync.WaitGroup),
		stopWaitTime: opt.StopWaitTime,
	}
	err := server.init()
	if err != nil {
		return nil, nil, err
	}

	return server, ctx, nil
}

func (s *ServerExample3) init() error {
	return nil
}

func (s *ServerExample3) Name() string {
	return s.name
}

func (s *ServerExample3) GetCtx() *servercontext.ServerContext {
	return s.ctx
}

func (s *ServerExample3) Run() {
	s.running = true
	defer func() {
		s.running = false
	}()

	s.wg = new(sync.WaitGroup)
	s.wg.Add(1)
	defer s.wg.Done()

MainCycle:
	for {
		logger.Warnf("Example3: I am running!")

		select {
		case <-s.ctx.Listen():
			logger.Warnf("Example3: I am stop!")
			break MainCycle
		case <-time.After(2 * time.Second):
			continue
		}
	}
}

func (s *ServerExample3) Stop() {
	s.ctx.StopTask()
	if s.wg != nil {
		wgchan := make(chan any)

		go func() {
			s.wg.Wait()
			close(wgchan)
		}()

		select {
		case <-time.After(s.stopWaitTime):
			logger.Errorf("%s - 退出清理超时... (%s)", s.name, strconvutils.TimeDurationToString(s.stopWaitTime))
		case <-wgchan:
			// pass
		}
	}
}

func (s *ServerExample3) IsRunning() bool {
	return s.running
}

func _test() {
	var a serverinterface.Server
	var b *ServerExample3

	a = b
	_ = a
}
