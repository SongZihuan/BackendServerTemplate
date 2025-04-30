// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example3

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/goutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"sync"
	"sync/atomic"
	"time"
)

type ServerExample3 struct {
	running      atomic.Bool
	ctx          *servercontext.ServerContext
	name         string
	wg           *sync.WaitGroup
	stopWaitTime time.Duration
	lockThread   bool
}

type ServerExample3Option struct {
	StopWaitTime time.Duration
	LockThread   bool
}

func NewServerExample3(opt *ServerExample3Option) (*ServerExample3, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ServerExample3Option{
			StopWaitTime: 10 * time.Second,
			LockThread:   false,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	server := &ServerExample3{
		ctx:          ctx,
		name:         "example3",
		wg:           new(sync.WaitGroup),
		stopWaitTime: opt.StopWaitTime,
		lockThread:   opt.LockThread,
	}

	server.running.Store(false)

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
	if s.running.Swap(true) {
		return
	}
	defer func() {
		s.running.Store(false)
	}()

	if s.lockThread {
		err := goutils.LockOSThread()
		if err != nil {
			s.ctx.RunError(err)
			return
		}
	}
	defer func() {
		if s.lockThread {
			_ = goutils.UnlockOSThread()
		}
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
	return s.running.Load()
}
