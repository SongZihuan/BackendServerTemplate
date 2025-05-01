// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configerror

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
)

type Error interface {
	Msg() string
	Error() string
	Warning() string
	IsError() bool
	IsWarning() bool
}

func NewErrorf(format string, args ...any) Error {
	msg := fmt.Sprintf(format, args...)

	logger.Errorf("config error: %s", msg)
	return &configError{msg: msg, isError: true}
}

func NewWarningf(format string, args ...any) Error {
	msg := fmt.Sprintf(format, args...)

	logger.Errorf("config warning: %s", msg)
	return &configError{msg: msg, isError: false}
}

func ShowWarningf(format string, args ...any) {
	_ = NewWarningf(format, args...)
}
