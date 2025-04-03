// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datefilewriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
	"os"
	"path"
	"time"
)

type DateFileWriter struct {
	dirPath        string
	filenamePrefix string
	filenameSuffix string
	file           *os.File
	close          bool
}

func (f *DateFileWriter) Write(p []byte) (n int, err error) {
	if f.close {
		return 0, fmt.Errorf("date file writer has been close")
	}

	suffix := time.Now().Format(time.DateOnly)
	if suffix != f.filenameSuffix {
		_ = f.closeFile()
		err := f.openFile(suffix)
		if err == nil {
			return 0, err
		}
	}

	if f.file == nil {
		return 0, fmt.Errorf("file writer has been close")
	}

	if f.file.Fd() == ^(uintptr(0)) { // 检查文件描述符是否为 -1
		return 0, fmt.Errorf("file writer has been close")
	}

	return f.file.Write(p)
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
	defer func() {
		f.filenameSuffix = newSuffix
	}()

	if f.file != nil {
		return fmt.Errorf("last file has not been closse")
	}

	filename := fmt.Sprintf("%s-%s", f.filenamePrefix, newSuffix)
	filePath := path.Join(f.dirPath, filename)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	f.file = file

	return nil
}

func (f *DateFileWriter) Close() error {
	return f.ExitClose()
}

func (f *DateFileWriter) ExitClose() error {
	defer func() {
		f.file = nil
	}()

	if f.file != nil {
		return f.file.Close()
	}

	return nil
}

func NewDateFileWriter(dirpath string, filenamePrefix string) (*DateFileWriter, error) {
	var writer write.WriteCloser
	var res = new(DateFileWriter)

	if filesystemutils.IsFile(dirpath) {
		return nil, fmt.Errorf("dir not exists")
	}

	err := os.MkdirAll(dirpath, 0644)
	if err != nil {
		return nil, err
	}

	res.filenamePrefix = filenamePrefix
	res.close = false

	writer = res // 用于检验StdWriter实现了io.WriteCloser
	_ = writer

	return res, nil
}
