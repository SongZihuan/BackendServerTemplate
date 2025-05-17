// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package warpwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/utils/osutils"
	"io"
	"sync"
)

type WarpWrite struct {
	level  loglevel.LoggerLevel
	tag    bool
	writer io.Writer
	close  bool // 表示 warp 关闭时，是否关闭 writer
	fn     logformat.FormatFunc
	mutex  sync.Mutex
}

func (w *WarpWrite) Write(data *logformat.LogData) chan any {
	res := make(chan any)

	// 此处 w.level 是只读的，因此可以不上锁操作
	if (w.level.Int() > data.Level.Int()) || (data.Level == loglevel.PseudoLevelTag && !w.tag) {
		close(res)
		return res
	}

	go func() {
		w.write(data)
		close(res)
	}()

	return res
}

func (w *WarpWrite) write(data *logformat.LogData) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.writer == nil {
		return
	}

	_, _ = fmt.Fprintf(w.writer, "%s\n", w.fn(data))

	syncer, ok := w.writer.(osutils.Syncer)
	if ok {
		_ = syncer.Sync()
	}
}

func (w *WarpWrite) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	defer func() {
		w.writer = nil
		w.close = false
	}()

	if closer, ok := w.writer.(io.WriteCloser); w.close && ok {
		return closer.Close()
	}

	return nil
}

func NewWarpWriter(level loglevel.LoggerLevel, tag bool, w io.Writer, fn logformat.FormatFunc) *WarpWrite {
	return &WarpWrite{
		level:  level,
		tag:    tag,
		writer: w,
		fn:     fn,
		close:  false,
	}
}

func NewWarpWriteCloser(level loglevel.LoggerLevel, tag bool, w io.Writer, fn logformat.FormatFunc) *WarpWrite {
	return &WarpWrite{
		level:  level,
		tag:    tag,
		writer: w,
		fn:     fn,
		close:  true,
	}
}
