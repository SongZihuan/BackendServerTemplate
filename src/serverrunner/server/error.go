// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import "time"

type RunStartupError interface {
	Error() error
	CleanTime() time.Duration
	Wait() chan any
}

type runStartupError struct {
	err       error
	wait      chan any
	cleanTime time.Duration
}

func NewRunStartupError(err error, wait chan any, cleanTime time.Duration) RunStartupError {
	return &runStartupError{
		err:       err,
		wait:      wait,
		cleanTime: cleanTime,
	}
}

func (r *runStartupError) Error() error {
	return r.err
}

func (r *runStartupError) CleanTime() time.Duration {
	return r.cleanTime
}

func (r *runStartupError) Wait() chan any {
	return r.wait
}
