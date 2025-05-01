// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/server"
	"github.com/SongZihuan/BackendServerTemplate/src/serverrunner/servercontext"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/strconvutils"
	"reflect"
	"sync"
	"time"
)

type child struct {
	server map[string]*server.Server
	ctx    map[string]*servercontext.ServerContext
	wg     sync.WaitGroup
}

type self struct {
	name   string
	server *server.Server
	ctx    *servercontext.ServerContext
}

type ControllerCore struct {
	child child
	self  self
}

type ControllerCoreOption struct{}

func NewControllerCore(opt *ControllerCoreOption) (*ControllerCore, error) {
	if opt == nil {
		opt = &ControllerCoreOption{}
	}

	controller := &ControllerCore{
		child: child{
			server: make(map[string]*server.Server, 10),
			ctx:    make(map[string]*servercontext.ServerContext, 10),
		},
		self: self{
			name: server.ControllerName,
		},
	}

	return controller, nil
}

func (cc *ControllerCore) Name() string {
	return cc.self.name
}

func (cc *ControllerCore) Init(s *server.Server, ctx *servercontext.ServerContext) error {
	cc.self.server = s
	cc.self.ctx = ctx

	cc.child.server[cc.self.name] = s
	cc.child.ctx[cc.self.name] = ctx

	return nil
}

func (cc *ControllerCore) AddServerCore(option *server.ServerOption) (*server.Server, error) {
	if cc.self.server.Status() != server.StatusWaitRun {
		return nil, fmt.Errorf("controller is not in wait run status")
	}

	ser, _, err := server.NewServer(option)
	if err != nil {
		return nil, err
	}

	return ser, cc.addServer(ser)
}

func (cc *ControllerCore) AddServer(ser *server.Server) error {
	if cc.self.server.Status() != server.StatusWaitRun {
		return fmt.Errorf("controller is not in wait run status")
	}

	return cc.addServer(ser)
}

func (cc *ControllerCore) addServer(ser *server.Server) error {
	if ser.IsController() {
		return fmt.Errorf("can not add controller to controller")
	}

	name := ser.Name()

	if _, ok1 := cc.child.server[name]; ok1 {
		return fmt.Errorf("server is exists")
	} else if _, ok2 := cc.child.ctx[name]; ok2 {
		logger.Panicf("the service does not exist, but the server %s context is retained and the data is incomplete.", ser.Name())
		return fmt.Errorf("server context is exists")
	}

	cc.child.server[name] = ser
	cc.child.ctx[name] = ser.GetCtx()

	return nil
}

func (cc *ControllerCore) DelServer(ser *server.Server) error {
	if cc.self.server.Status() != server.StatusWaitRun {
		return fmt.Errorf("controller is not in wait run status")
	}

	name := ser.Name()

	if _, ok1 := cc.child.server[name]; !ok1 {
		return fmt.Errorf("server is not exists")
	} else if _, ok2 := cc.child.ctx[name]; !ok2 {
		logger.Panicf("controller record error: the server %s exists, but the context is lost and the data is incomplete.", ser.Name())
		return fmt.Errorf("server context is not exists")
	}

	delete(cc.child.server, name)
	delete(cc.child.server, name)

	return nil
}

func (cc *ControllerCore) Run() error {
	for name, ser := range cc.child.server {
		if name == cc.self.name {
			continue
		}

		_, ok := cc.child.ctx[name]
		if !ok {
			logger.Errorf("server %s context not found", name)
			continue
		}

		logger.Infof("start to run server: %s", name)
		err, timeout := server.RunWithWorkGroup(ser, &cc.child.wg)
		if err != nil {
			logger.Errorf("start server %s error: %s", ser.Name(), err.Error())
		} else if timeout {
			logger.Warnf("start server %s by %s success. but check timeout", ser.Name(), cc.self.name)
		} else {
			logger.Warnf("start server %s by %s success", ser.Name(), cc.self.name)
		}
	}

MainCycle:
	for {
		if len(cc.child.ctx) == 0 {
			break
		} else if _, ok := cc.child.ctx[cc.self.name]; ok && len(cc.child.ctx) == 1 {
			break MainCycle
		}

		cases := make([]reflect.SelectCase, 0, len(cc.child.ctx))
		serverName := make([]string, 0, len(cc.child.ctx))

		for name, ctx := range cc.child.ctx {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ctx.Listen()),
			})
			serverName = append(serverName, name)
		}

		chosen, _, _ := reflect.Select(cases)
		stopServerName := serverName[chosen]

		ctx, ok1 := cc.child.ctx[stopServerName]
		if !ok1 {
			logger.Panicf("unknown server stop: %s", stopServerName)
			continue MainCycle
		} else if _, ok2 := cc.child.server[stopServerName]; !ok2 {
			logger.Panicf("unknown server stop: %s", stopServerName)
			continue MainCycle
		}

		if !ctx.IsStop() {
			logger.Errorf("server %s exit false positive", stopServerName)
			continue MainCycle
		}

		if stopServerName == cc.self.name {
			delete(cc.child.ctx, stopServerName)
			delete(cc.child.server, stopServerName)
			break MainCycle
		}

		switch ctx.Reason() {
		default:
			logger.Panicf("error server context reason: %d", ctx.Reason())
			continue MainCycle
		case servercontext.StopReasonStop:
			err := ctx.Error()
			if err == nil {
				logger.Infof("server %s stop", stopServerName)
			} else {
				logger.Infof("server %s stop with error: %s", stopServerName, err.Error())
			}
		case servercontext.StopReasonStopAllTask:
			err := ctx.Error()
			if err == nil {
				logger.Infof("server %s stop (all task stop)", stopServerName)
			} else {
				logger.Infof("server %s stop with error: %s (all task stop)", stopServerName, err.Error())
			}
			cc.self.ctx.FinishErrorAndStopAllTask(err) // 会自动判断err是否为nil
		case servercontext.StopReasonFinish:
			// 不会停止其他任务
			err := ctx.Error()
			if err == nil {
				logger.Infof("server %s finish", stopServerName)
			} else {
				logger.Infof("server %s finish with error: %s", stopServerName, err.Error())
			}
		case servercontext.StopReasonFinishAndStopAllTask:
			err := ctx.Error()
			if err == nil {
				logger.Infof("server %s finish (all task stop)", stopServerName)
			} else {
				logger.Infof("server %s finish with error: %s (all task stop)", stopServerName, err.Error())
			}
			cc.self.ctx.FinishErrorAndStopAllTask(err) // 会自动判断err是否为nil
		}

		delete(cc.child.ctx, stopServerName)
		delete(cc.child.server, stopServerName)
		continue MainCycle
	}

	return nil
}

func (cc *ControllerCore) Stop() {
	go func() {
		for _, ser := range cc.child.server {
			ser.Stop()
		}
	}()

	go func() {
		wgchan := make(chan any)

		go func() {
			cc.child.wg.Wait()
			close(wgchan)
		}()

		select {
		case <-time.After(cc.GetStopWaitTime()):
			logger.Errorf("%s - Exit subtask timeout... (%s)", cc.self.name, strconvutils.TimeDurationToString(cc.GetStopWaitTime()))
		case <-wgchan:
			logger.Warnf("%s - Exit subtask completed", cc.self.name)
		}
	}()
}

func (cc *ControllerCore) Controller() {}

func (cc *ControllerCore) GetStopWaitTime() time.Duration {
	stopWaitTime := time.Duration(0) // 自己也作为记录的一部分
	for name, ser := range cc.child.server {
		if name == cc.self.name {
			continue
		}
		stopWaitTime = max(stopWaitTime, ser.StopWaitTime())
	}
	return stopWaitTime
}

func (cc *ControllerCore) GetStartupWaitTime() time.Duration {
	startupWaitTime := time.Duration(0)
	for name, ser := range cc.child.server {
		if name == cc.self.name {
			continue
		}
		startupWaitTime = max(startupWaitTime, ser.StartupWaitTime())
	}
	return startupWaitTime
}
