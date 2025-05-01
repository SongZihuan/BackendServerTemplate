// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filewriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"os"
)

type FileWriter struct {
	filePath string
	file     *os.File
	fn       logformat.FormatFunc
}

func (f *FileWriter) Write(data *logformat.LogData) (n int, err error) {
	if !fileutils.IsFileOpen(f.file) {
		return 0, fmt.Errorf("file writer has been close")
	}

	return fmt.Fprintf(f.file, "%s\n", f.fn(data))
}

func (f *FileWriter) Close() error {
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

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	res.filePath = filepath
	res.file = file
	res.fn = fn

	return res, nil
}
