// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datefilewriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/fileutils"
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
	fn             logformat.FormatFunc
}

func (f *DateFileWriter) Write(data *logformat.LogData) (n int, err error) {
	if f.close {
		return 0, fmt.Errorf("date file writer has been close")
	}

	suffix := time.Now().Format(time.DateOnly)
	if suffix != f.filenameSuffix {
		_ = f.closeFile()
		err := f.openFile(suffix)
		if err != nil {
			return 0, err
		}
	}

	if fileutils.IsFileOpen(f.file) {
		return 0, fmt.Errorf("file writer has been close")
	}

	return fmt.Fprintf(f.file, "%s\n", f.fn(data))
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

	filename := fmt.Sprintf("%s.%s.log", f.filenamePrefix, newSuffix)
	filePath := path.Join(f.dirPath, filename)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.file = file
	f.filenameSuffix = newSuffix

	return nil
}

func (f *DateFileWriter) Close() error {
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

func _testDateFileWriter() {
	var a write.WriteCloser
	var b *DateFileWriter

	a = b
	_ = a
}
