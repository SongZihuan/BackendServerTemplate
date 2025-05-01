// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example2

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/config"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/goutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"sync"
	"sync/atomic"
	"time"
)

type ServerExample2 struct {
	running      atomic.Bool
	ctx          *servercontext.ServerContext
	name         string
	wg           *sync.WaitGroup
	stopWaitTime time.Duration
	lockThread   bool
}

type ServerExample2Option struct {
	StopWaitTime time.Duration
	LockThread   bool
}

func NewServerExample2(opt *ServerExample2Option) (*ServerExample2, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ServerExample2Option{
			StopWaitTime: 10 * time.Second,
			LockThread:   false,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	server := &ServerExample2{
		ctx:          ctx,
		name:         "example2",
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

func (s *ServerExample2) init() error {
	return nil
}

func (s *ServerExample2) Name() string {
	return s.name
}

func (s *ServerExample2) GetCtx() *servercontext.ServerContext {
	return s.ctx
}

func (s *ServerExample2) Run() {
	if s.running.Swap(true) {
		return
	}
	defer func() {
		s.running.Store(false)
	}()

	if s.lockThread {
		err := goutils.LockOSThread()
		if err != nil {
			s.ctx.FinishError(err)
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
		fmt.Printf("Example2: Hello, %s. I am running!\n", config.Data().Server.Name)

		select {
		case <-s.ctx.Listen():
			fmt.Println("Example2: I am stop!")
			break MainCycle
		case <-time.After(2 * time.Second):
			continue
		}
	}
}

func (s *ServerExample2) Stop() {
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

func (s *ServerExample2) IsRunning() bool {
	return s.running.Load()
}
