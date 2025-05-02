// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datefilewriter

import (
	"context"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/gofrs/flock"
	"os"
	"path"
	"sync"
	"time"
)

type DateFileWriter struct {
	dirPath        string
	filenamePrefix string
	filenameSuffix string
	fileName       string
	filePath       string
	file           *os.File
	fileLockPath   string
	fileLock       *flock.Flock
	close          bool
	fn             logformat.FormatFunc
	lock           sync.Mutex
}

func (f *DateFileWriter) Write(data *logformat.LogData) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.close {
		return
	}

	suffix := time.Now().Format(time.DateOnly)
	if suffix != f.filenameSuffix {
		_ = f.closeFile()
		err := f.openFile(suffix)
		if err != nil {
			return
		}
	}

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
}

func (f *DateFileWriter) closeFile() error {
	defer func() {
		f.file = nil
	}()

	if f.file != nil {
		return f.file.Close()
	}

	return nil
}

func (f *DateFileWriter) openFile(newSuffix string) error {
	if f.file != nil {
		return fmt.Errorf("last file has not been closse")
	}

	f.fileName = fmt.Sprintf("%s.%s.log", f.filenamePrefix, newSuffix)
	f.filePath = path.Join(f.dirPath, f.fileName)
	f.fileLockPath = f.filePath + ".lock"

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return err
	}

	f.file = file
	f.filenameSuffix = newSuffix
	f.fileLock = flock.New(f.fileLockPath)

	return nil
}

func (f *DateFileWriter) Close() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	defer func() {
		f.file = nil
		f.close = true
	}()

	if f.file != nil {
		return f.file.Close()
	}

	return nil
}

func NewDateFileWriter(dirpath string, filenamePrefix string, fn logformat.FormatFunc) (*DateFileWriter, error) {
	var writer write.WriteCloser
	var res = new(DateFileWriter)

	if filesystemutils.IsFile(dirpath) {
		return nil, fmt.Errorf("dir not exists")
	}

	err := os.MkdirAll(dirpath, 0644)
	if err != nil {
		return nil, err
	}

	res.dirPath = dirpath
	res.filenamePrefix = filenamePrefix
	res.close = false
	res.fn = fn

	writer = res // 用于检验StdWriter实现了io.WriteCloser
	_ = writer

	return res, nil
}
