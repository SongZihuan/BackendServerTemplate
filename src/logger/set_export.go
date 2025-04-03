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

func SetArgs0Name(args0 string, args0Name string) error {
	if !internal.IsReady() {
		return fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetArgs0Name(args0, args0Name)
}

func SetWarnWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetWarnWriter(w)
}

func SetErrWriter(w io.Writer) (io.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetErrWriter(w)
}
