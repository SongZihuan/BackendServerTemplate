// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/server/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/server/serverinterface"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type Controller struct {
	server       map[string]serverinterface.Server
	context      map[string]*servercontext.ServerContext
	ctx          *servercontext.ServerContext
	running      atomic.Bool
	name         string
	stopWaitTime time.Duration
	wg           *sync.WaitGroup
}

type ControllerOption struct {
	StopWaitTime time.Duration
}

func NewController(opt *ControllerOption) (*Controller, error) {
	ctx := servercontext.NewServerContext()

	if opt == nil {
		opt = &ControllerOption{
			StopWaitTime: 10 * time.Second,
		}
	} else {
		if opt.StopWaitTime == 0 {
			opt.StopWaitTime = 10 * time.Second
		}
	}

	controller := &Controller{
		server:       make(map[string]serverinterface.Server, 10),
		context:      make(map[string]*servercontext.ServerContext, 10),
		ctx:          ctx,
		name:         serverinterface.ControllerName,
		wg:           new(sync.WaitGroup),
		stopWaitTime: opt.StopWaitTime,
	}

	{
		controller.server[controller.name] = controller
		controller.context[controller.name] = ctx
	}

	return controller, nil
}

func (s *Controller) AddServer(ser serverinterface.Server) error {
	if s.running.Load() {
		return fmt.Errorf("controller is running")
	}

	name := ser.Name()
	if name == serverinterface.ControllerName {
		return fmt.Errorf("name can not be %s", serverinterface.ControllerName)
	} else if name == "" {
		return fmt.Errorf("name can not be empty")
	}

	_, ok := s.server[name]
	if ok {
		return fmt.Errorf("server is exists")
	}

	s.server[name] = ser
	s.context[name] = ser.GetCtx()

	return nil
}

func (s *Controller) DelServer(ser serverinterface.Server) error {
	if s.running.Load() {
		return fmt.Errorf("controller is running")
	}

	name := ser.Name()
	if name == serverinterface.ControllerName {
		return fmt.Errorf("name can not be %s", serverinterface.ControllerName)
	} else if name == "" {
		return fmt.Errorf("name can not be empty")
	}

	_, ok := s.server[name]
	if !ok {
		return fmt.Errorf("server is not exists")
	}

	delete(s.server, name)
	delete(s.server, name)

	return nil
}

func (s *Controller) Name() string {
	return s.name
}

func (s *Controller) GetCtx() *servercontext.ServerContext {
	return s.ctx
}

func (s *Controller) Run() {
	if s.running.Swap(true) {
		return
	}
	defer func() {
		s.running.Store(false)
	}()

	s.wg = new(sync.WaitGroup)
	s.wg.Add(1)
	defer s.wg.Done()

	for name, server := range s.server {
		if name == s.name {
			continue
		}

		_, ok := s.context[name]
		if !ok {
			logger.Errorf("server %s context not found", name)
			continue
		}

		logger.Infof("start to run server: %s", name)

		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			server.Run()
		}()
	}

SelectCycle:
	for {
		cases := make([]reflect.SelectCase, 0, len(s.context))
		serverName := make([]string, 0, len(s.context))

		for name, ctx := range s.context {
			if name == s.name {
				continue
			}

			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ctx.Listen()),
			})
			serverName = append(serverName, name)
		}

		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(s.ctx.Listen()),
		})
		serverName = append(serverName, s.name)

		chosen, _, _ := reflect.Select(cases)
		stopServerName := serverName[chosen]

		if stopServerName != s.name {
			ctx, ok := s.context[stopServerName]
			if !ok {
				logger.Errorf("unknown server stop: %s", stopServerName)
				break SelectCycle
			} else {
				switch ctx.Reason() {
				default:
					fallthrough
				case servercontext.StopReasonStop:
					logger.Infof("server %s stop", stopServerName)
					break SelectCycle
				case servercontext.StopReasonFinish:
					// 不会停止其他任务
					logger.Infof("server %s run finished", stopServerName)
					delete(s.context, stopServerName)
					delete(s.server, stopServerName)
					continue SelectCycle
				case servercontext.StopReasonError:
					err := ctx.Error()
					if errors.Is(err, servercontext.StopAllTask) {
						logger.Infof("server %s run finished (stop all task)", stopServerName)
						s.ctx.RunError(err)
					} else if err != nil {
						logger.Infof("server %s stop with error: %s", stopServerName, err.Error())
						s.ctx.FinishAndStopAllTask()
					} else {
						logger.Infof("server %s stop with error: unknown", stopServerName)
						s.ctx.RunError(err)
					}
					break SelectCycle
				}
			}
		} else {
			break SelectCycle
		}
	}

	for name, ctx := range s.context {
		if name == s.name {
			continue
		}

		ctx.StopTask()
	}
}

func (s *Controller) Stop() {
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

func (s *Controller) IsRunning() bool {
	return s.running.Load()
}

func _test() {
	var a serverinterface.Server
	var b *Controller

	a = b
	_ = a
}
