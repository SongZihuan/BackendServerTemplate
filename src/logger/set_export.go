// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"io"
)

func SetLevel(level loglevel.LoggerLevel) error {
	if !internal.IsReady() {
		return fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetLevel(level)
}

func SetLogTag(logTag bool) error {
	if !internal.IsReady() {
		return fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetLogTag(logTag)
}

func SetHumanWarnWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetHumanWarnWriter(w)
}

func SetHumanErrWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetHumanErrWriter(w)
}

func SetMachineWarnWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetMachineWarnWriter(w)
}

func SetMachineErrWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetMachineErrWriter(w)
}
