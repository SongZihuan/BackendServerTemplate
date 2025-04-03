// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package servercontext

import (
	"fmt"
	"sync"
)

type StopReason int

var StopAllTask = fmt.Errorf("stop all task")

const (
	StopReasonStop StopReason = iota + 1
	StopReasonFinish
	StopReasonError
)

type ServerContext struct {
	mutex    sync.Mutex
	stopchan chan any
	err      error
	reason   StopReason
}

func NewServerContext() *ServerContext {
	return &ServerContext{
		stopchan: make(chan any),
		err:      nil,
		reason:   0,
	}
}

func (c *ServerContext) Listen() <-chan any {
	return c.stopchan
}

// StopTask 表示外部环境终止任务
func (c *ServerContext) StopTask() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.reason = StopReasonStop
	close(c.stopchan)
}

// Finish 表示任务内部环境完成任务
func (c *ServerContext) Finish() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.reason = StopReasonFinish
	close(c.stopchan)
}

// FinishAndStopAllTask 表示任务内部环境完成任务，并且退出所有任务
func (c *ServerContext) FinishAndStopAllTask() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.err = StopAllTask
	c.reason = StopReasonError
	close(c.stopchan)
}

// RunError 表示任务内部环境运行遇到错误
func (c *ServerContext) RunError(err error) {
	if err == nil {
		c.Finish()
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.err = err
	c.reason = StopReasonError
	close(c.stopchan)
}

func (c *ServerContext) IsStop() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.isStop()
}

func (c *ServerContext) Error() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.err
}

func (c *ServerContext) Reason() StopReason {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.reason
}

func (c *ServerContext) isStop() bool {
	select {
	case <-c.stopchan:
		return true
	default:
		return false
	}
}
