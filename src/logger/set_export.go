// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
)

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
