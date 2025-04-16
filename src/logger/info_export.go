// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger/internal"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"io"
)

func GetLevel() loglevel.LoggerLevel {
	if !internal.IsReady() {
		return loglevel.LevelDebug
	}
	return internal.GlobalLogger.GetLevel()
}

func IsLogTag() bool {
	if !internal.IsReady() {
		return false
	}
	return internal.GlobalLogger.IsLogTag()
}

func GetWarnWriter() io.Writer {
	if !internal.IsReady() {
		return nil
	}
	return internal.GlobalLogger.GetWarnWriter()
}

func GetErrWriter() io.Writer {
	if !internal.IsReady() {
		return nil
	}
	return internal.GlobalLogger.GetErrWriter()
}

func IsWarnWriterTerm() bool {
	if !internal.IsReady() {
		return false
	}

	return internal.GlobalLogger.IsWarnWriterTerm()
}

func IsErrWriterTerm() bool {
	if !internal.IsReady() {
		return false
	}

	return internal.GlobalLogger.IsErrWriterTerm()
}

func IsTermDump() bool {
	if !internal.IsReady() {
		return false
	}

	return internal.GlobalLogger.IsTermDump()
}

func IsWarnWriterTermNoDump() bool {
	if !internal.IsReady() {
		return false
	}

	return internal.GlobalLogger.IsWarnWriterTermNoDump()
}

func IsErrWriterTermNoDump() bool {
	if !internal.IsReady() {
		return false
	}

	return internal.GlobalLogger.IsErrWriterTermNoDump()
}
