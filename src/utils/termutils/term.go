// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package termutils

import (
	"github.com/mattn/go-isatty"
	"io"
	"os"
)

func IsTerm(writer io.Writer) bool {
	w, ok := writer.(*os.File)
	if !ok {
		return false
	}

	if !isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()) { // 非终端
		return false
	}

	return true
}

func EnvHasTermDump() bool {
	// TERM为dump表示终端为基础模式，不支持高级显示
	return os.Getenv("TERM") == "dumb"
}

func IsTermDump(writer io.Writer) bool {
	return EnvHasTermDump() && IsTerm(writer)
}

func IsTermAdvanced(writer io.Writer) bool {
	return !EnvHasTermDump() && IsTerm(writer)
}
