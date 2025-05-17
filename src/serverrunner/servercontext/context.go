// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package servercontext

import (
	"sync"
)

type StopReason int

const (
	StopReasonStop                 StopReason = iota + 1 // 外部停止
	StopReasonStopAllTask                                // 外部停止，并请求 controller 也停止
	StopReasonFinish                                     // 内部停止
	StopReasonFinishAndStopAllTask                       // 内部停止，并请求 controller 也停止
)

type ServerContext struct {
	mutex    sync.Mutex
	stopchan chan any   // 控制状态（若为close则表示停止运行）
	err      error      // 运行错误
	reason   StopReason // 停止原因
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

// StopTaskError 表示外部环境终止任务并伴随错误
func (c *ServerContext) StopTaskError(err error) {
	if err == nil {
		c.StopTask()
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.reason = StopReasonStop
	c.err = err
	close(c.stopchan)
}

// StopAllTask 表示外部环境终止全部任务
func (c *ServerContext) StopAllTask() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.reason = StopReasonStopAllTask
	close(c.stopchan)
}

// StopAllTaskError 表示外部环境终止全部任务并伴随错误
func (c *ServerContext) StopAllTaskError(err error) {
	if err == nil {
		c.StopAllTask()
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.reason = StopReasonStopAllTask
	c.err = err
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

	c.reason = StopReasonFinishAndStopAllTask
	close(c.stopchan)
}

// FinishError 表示任务内部环境运行遇到错误
func (c *ServerContext) FinishError(err error) {
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
	c.reason = StopReasonFinish
	close(c.stopchan)
}

// FinishErrorAndStopAllTask 表示任务内部环境运行遇到错误，并退出全部服务
func (c *ServerContext) FinishErrorAndStopAllTask(err error) {
	if err == nil {
		c.FinishAndStopAllTask()
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop() {
		return
	}

	c.err = err
	c.reason = StopReasonFinishAndStopAllTask
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
