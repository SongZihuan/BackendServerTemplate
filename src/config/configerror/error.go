// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configerror

type configError struct {
	msg     string
	isError bool
}

func (e *configError) Msg() string {
	if e.isError {
		return e.Error()
	}
	return e.Warning()
}

func (e *configError) Error() string {
	return e.msg
}

func (e *configError) Warning() string {
	return e.msg
}

func (e *configError) IsError() bool {
	return e.isError
}

func (e *configError) IsWarning() bool {
	return !e.isError
}
