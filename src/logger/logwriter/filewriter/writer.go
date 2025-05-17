// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filewriter

import (
	"context"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/gofrs/flock"
	"os"
	"path"
	"sync"
	"time"
)

type FileWriter struct {
	level        loglevel.LoggerLevel
	tag          bool
	filePath     string
	file         *os.File
	fileLockPath string
	fileLock     *flock.Flock
	fn           logformat.FormatFunc
	mutex        sync.Mutex
}

func (w *FileWriter) Write(data *logformat.LogData) chan any {
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

func (f *FileWriter) write(data *logformat.LogData) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if !fileutils.IsFileOpen(f.file) {
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	ok, err := f.fileLock.TryLockContext(ctx, 1*time.Second)
	if err != nil || !ok {
		return
	}
	defer func() {
		_ = f.fileLock.Unlock()
		_ = os.Remove(f.fileLockPath)
	}()

	_, _ = fmt.Fprintf(f.file, "%s\n", f.fn(data))

	return
}

func (f *FileWriter) Close() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.file == nil {
		return nil
	}

	defer func() {
		f.file = nil
	}()

	return f.file.Close()
}

func NewFileWriter(level loglevel.LoggerLevel, tag bool, filepath string, fn logformat.FormatFunc) (*FileWriter, error) {
	var res = new(FileWriter)

	if fn == nil {
		fn = logformat.FormatFile
	}

	dir := path.Dir(filepath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return nil, err
	}

	res.level = level
	res.tag = tag
	res.filePath = filepath
	res.fileLockPath = res.filePath + ".lock"
	res.file = file
	res.fileLock = flock.New(res.fileLockPath)
	res.fn = fn

	return res, nil
}
