// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package example1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/goutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"sync"
	"sync/atomic"
	"time"
)

type ServerExample1 struct {
	running      atomic.Bool
	ctx          *servercontext.ServerContext
	name         string
	wg           *sync.WaitGroup
	stopWaitTime time.Duration
	lockThread   bool
}

type ServerExample1Option struct {
	StopWaitTime time.Duration
	LockThread   bool
}

func NewServerExample1(opt *ServerExample1Option) (*ServerExample1, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ServerExample1Option{
			StopWaitTime: 10 * time.Second,
			LockThread:   false,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	server := &ServerExample1{
		ctx:          ctx,
		name:         "example1",
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
		if global.GitTag == "" || global.GitTagCommitHash == "" {
			fmt.Printf("Example1: I am running! BuildDate: '%s' Commit: '%s' Version: '%s' Now: '%s'\n", global.BuildTime.Format(time.DateTime), global.GitCommitHash, global.Version, time.Now().Format(time.DateTime))
		} else {
			fmt.Printf("Example1: I am running! BuildDate: '%s' Commit: '%s' Tag: '%s' Tag Commit: '%s' Version: '%s' Now: '%s'\n", global.BuildTime.Format(time.DateTime), global.GitCommitHash, global.GitTag, global.GitTagCommitHash, global.Version, time.Now().Format(time.DateTime))
		}

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
			logger.Errorf("%s - 退出清理超时... (%s)", s.name, strconvutils.TimeDurationToString(s.stopWaitTime))
		case <-wgchan:
			// pass
		}
	}
}

func (s *ServerExample1) IsRunning() bool {
	return s.running.Load()
}
