// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stdutils

import (
	"fmt"
	"os"
)

var (
	bakStdin  = os.Stdin
	bakStdout = os.Stdout
	bakStderr = os.Stderr
)

var nullFile *os.File = nil

func QuiteMode() error {
	if nullFile == nil {
		err := OpenNullFile()
		if err != nil {
			return err
		}
	}

	os.Stdin = nullFile
	os.Stdout = nullFile
	os.Stderr = nullFile

	return nil
}

func OpenNullFile() error {
	if nullFile != nil {
		return fmt.Errorf("null file is exists")
	}

	_nullFile, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return err
	}

	nullFile = _nullFile
	return nil
}

func Recover() {
	os.Stdin = bakStdin
	os.Stdout = bakStdout
	os.Stderr = bakStderr
}

func CloseNullFile() {
	if nullFile == nil {
		return
	}
	_ = nullFile.Close()
}
