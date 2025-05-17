// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/utils/goutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/strconvutils"
	"sync"
	"sync/atomic"
	"time"
)

const (
	StatusWaitInit int32 = iota
	StatusWaitRun
	StatusRunning
	StatusStopping
	StatusStop
)

type Server struct {
	name   string
	status atomic.Int32
	ctx    *servercontext.ServerContext
	wg     sync.WaitGroup

	lockThread     bool
	core           ServerCore
	controllerCore ControllerServerCore

	stopWaitTimeUseSpecifiedValue    bool // 仅对controller有效 （若为 false 则实际的 stopWaitTime 根据子服务的最大值和本 controller 的 stopWaitTime 共同决定）
	stopWaitTime                     time.Duration
	startupWaitTimeUseSpecifiedValue bool // 仅对于controller有效 （若为false 则实际的 startupWaitTime 取值为本 controller 设置的 startupWaitTime）
	startupWaitTime                  time.Duration
}

type ServerOption struct {
	StopWaitTimeUseSpecifiedValue    bool // 仅对controller有效
	StopWaitTime                     time.Duration
	StartupWaitTimeUseSpecifiedValue bool // 仅对于controller有效
	StartupWaitTime                  time.Duration
	LockThread                       bool
	ServerCore                       ServerCore
}

func NewServer(opt *ServerOption) (*Server, *servercontext.ServerContext, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		return nil, nil, fmt.Errorf("option is nil")
	}

	if opt.ServerCore == nil {
		return nil, nil, fmt.Errorf("server core is nil")
	}

	name := opt.ServerCore.Name()
	if name == "" {
		return nil, nil, fmt.Errorf("name is empty")
	}

	controllerCore, ok := opt.ServerCore.(ControllerServerCore)
	if ok && name != ControllerName {
		return nil, nil, fmt.Errorf("the name of controller server error: must be '%s' but '%s'", ControllerName, name)
	} else if !ok && name == ControllerName {
		return nil, nil, fmt.Errorf("the name of server (not controller) error: can not be '%s'", ControllerName)
	}

	server := &Server{
		name:                             name,
		ctx:                              ctx,
		lockThread:                       opt.LockThread,
		core:                             opt.ServerCore,
		controllerCore:                   controllerCore,
		stopWaitTimeUseSpecifiedValue:    opt.StopWaitTimeUseSpecifiedValue,
		stopWaitTime:                     opt.StopWaitTime,
		startupWaitTimeUseSpecifiedValue: opt.StartupWaitTimeUseSpecifiedValue,
		startupWaitTime:                  opt.StartupWaitTime,
	}

	server.status.Store(StatusWaitInit)

	err := server.init()
	if err != nil {
		return nil, nil, err
	}

	server.status.Store(StatusWaitRun)

	return server, ctx, nil
}

func (s *Server) init() error {
	return s.core.Init(s, s.ctx)
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) GetCtx() *servercontext.ServerContext {
	return s.ctx
}

func (s *Server) Run(startupErr chan RunStartupError) {
	defer logger.Recover()

	startupErrWait := make(chan any)
	defer close(startupErrWait)

	if !s.status.CompareAndSwap(StatusWaitRun, StatusRunning) {
		err := fmt.Errorf("server %s start run error: bad status %d", s.name, s.status.Load())
		logger.Errorf("%s", err.Error())
		startupErr <- NewRunStartupError(err, startupErrWait, 1*time.Second)
		close(startupErr)
		return
	}
	defer func() {
		if !s.status.CompareAndSwap(StatusStopping, StatusStop) {
			status := s.status.Swap(StatusStop) // 强行设置
			logger.Errorf("set server %s stop finish flag error: bad status %d", s.name, status)
		}
	}()

	s.wg.Add(1)
	defer s.wg.Done()

	if s.lockThread {
		err := goutils.LockOSThread()
		if err != nil {
			logger.Errorf("server %s lock os thread error: %s", s.name, err.Error())
			startupErr <- NewRunStartupError(err, startupErrWait, 1*time.Second)
			close(startupErr)
			return
		}
	}
	defer func() {
		if s.lockThread {
			err := goutils.UnlockOSThread()
			if err != nil {
				logger.Errorf("server %s unlock os thread error: %s", s.name, err.Error())
			}
		}
	}()

	defer func() {
		s.core.Stop()
	}()

	defer func() {
		if !s.status.CompareAndSwap(StatusRunning, StatusStopping) {
			status := s.status.Swap(StatusStopping) // 强行设置
			logger.Errorf("server %s stop run error: bad status %d", s.name, status)
		}
	}()

	close(startupErr)
	coreRunErr := s.core.Run()
	s.ctx.FinishError(coreRunErr) // 会自动判断 err != nil, 并且当 s.core.Run() 里面已经设置了退出，此处将会忽略
}

func (s *Server) StopAndWait() {
	status := s.status.Load()
	if status != StatusRunning && status != StatusStopping {
		return
	}

	s.stop()
	s.wait()
}

func (s *Server) Stop() {
	status := s.status.Load()
	if status != StatusRunning && status != StatusStopping {
		return
	}

	s.ctx.StopTask()
}

func (s *Server) stop() {
	s.ctx.StopTask()
}

func (s *Server) Wait() {
	status := s.status.Load()
	if status != StatusRunning && status != StatusStopping {
		return
	}

	s.wait()
}

func (s *Server) wait() {
	wgchan := make(chan any)

	go func() {
		s.wg.Wait()
		close(wgchan)
	}()

	select {
	case <-time.After(s.stopWaitTime):
		logger.Errorf("%s - Task timeout closed... (%s)", s.name, strconvutils.TimeDurationToString(s.stopWaitTime))
	case <-wgchan:
		logger.Warnf("%s - Task exit completed... ", s.name)
	}
}

func (s *Server) Status() int32 {
	return s.status.Load()
}

func (s *Server) IsController() bool {
	return s.controllerCore != nil
}

func (s *Server) AddServerCore(option *ServerOption) (*Server, error) {
	if s.controllerCore == nil {
		return nil, fmt.Errorf("not a controller server")
	}

	if s.Status() != StatusWaitRun {
		return nil, fmt.Errorf("controller is not in wait init status")
	}

	ser, err := s.controllerCore.AddServerCore(option)
	if err != nil {
		return ser, nil // ser也要同步返回，ser可能已经被创建
	}

	return ser, nil
}

func (s *Server) AddServer(ser *Server) error {
	if s.controllerCore == nil {
		return fmt.Errorf("not a controller server")
	}

	if s.Status() != StatusWaitRun {
		return fmt.Errorf("controller is not in wait init status")
	}

	err := s.controllerCore.AddServer(ser)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) DelServer(ser *Server) error {
	if s.controllerCore == nil {
		return fmt.Errorf("not a controller server")
	}

	if s.Status() != StatusWaitRun {
		return fmt.Errorf("controller is not in wait init status")
	}

	err := s.controllerCore.DelServer(ser)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) StartupWaitTime() time.Duration {
	if s.IsController() && s.startupWaitTimeUseSpecifiedValue {
		return s.controllerCore.GetStartupWaitTime() + s.startupWaitTime
	}
	return s.startupWaitTime
}

func (s *Server) StopWaitTime() time.Duration {
	if s.IsController() && s.startupWaitTimeUseSpecifiedValue {
		return s.controllerCore.GetStopWaitTime() + s.stopWaitTime
	}
	return s.stopWaitTime
}

func (s *Server) Runner() {}
