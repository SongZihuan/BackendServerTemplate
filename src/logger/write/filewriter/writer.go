// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filewriter

import (
	"context"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/gofrs/flock"
	"os"
	"sync"
	"time"
)

type FileWriter struct {
	filePath     string
	file         *os.File
	fileLockPath string
	fileLock     *flock.Flock
	fn           logformat.FormatFunc
	lock         sync.Mutex
}

func (f *FileWriter) Write(data *logformat.LogData) {
	f.lock.Lock()
	defer f.lock.Unlock()

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
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.file == nil {
		return nil
	}

	defer func() {
		f.file = nil
	}()

	return f.file.Close()
}

func NewFileWriter(filepath string, fn logformat.FormatFunc) (*FileWriter, error) {
	var res = new(FileWriter)

	if fn == nil {
		fn = logformat.FormatFile
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return nil, err
	}

	res.filePath = filepath
	res.fileLockPath = res.filePath + ".lock"
	res.file = file
	res.fileLock = flock.New(res.fileLockPath)
	res.fn = fn

	return res, nil
}
