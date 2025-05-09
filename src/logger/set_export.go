// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
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

func SetWarnWriter(w logwriter.Writer) (logwriter.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetWarnWriter(w)
}

func SetErrWriter(w logwriter.Writer) (logwriter.Writer, error) {
	if !internal.IsReady() {
		return nil, fmt.Errorf("logger not ready")
	}
	return internal.GlobalLogger.SetErrWriter(w)
}
