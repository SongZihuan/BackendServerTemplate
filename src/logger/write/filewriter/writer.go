// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filewriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"os"
)

type FileWriter struct {
	filePath string
	file     *os.File
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	if f.file == nil {
		return 0, fmt.Errorf("file writer has been close")
	}

	if f.file.Fd() == ^(uintptr(0)) { // 检查文件描述符是否为 -1
		return 0, fmt.Errorf("file writer has been close")
	}

	return f.file.Write(p)
}

func (f *FileWriter) Close() error {
	return f.ExitClose()
}

func (f *FileWriter) ExitClose() error {
	defer func() {
		f.file = nil
	}()

	if f.file != nil {
		return f.file.Close()
	}

	return nil
}

func NewFileWriter(filepath string) (*FileWriter, error) {
	var res = new(FileWriter)

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	res.filePath = filepath
	res.file = file

	return res, nil
}

func _testFileWriter() {
	var a write.WriteCloser
	var b *FileWriter

	a = b
	_ = a
}
