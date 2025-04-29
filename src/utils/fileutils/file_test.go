// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fileutils

import (
	"os"
	"path"
	"testing"
)

func TestFileIsOpen(t *testing.T) {
	temp, err := os.MkdirTemp("", "test*")
	if err != nil {
		t.Fatalf("create temp directory failed: %s", temp)
	}
	defer func() {
		_ = os.RemoveAll(temp)
	}()

	fileClose, err := os.OpenFile(path.Join(temp, "test_file_close"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		t.Fatalf("create temp/test_file failed: %s", temp)
	}
	_ = fileClose.Close()

	fileOpen, err := os.OpenFile(path.Join(temp, "test_file_open"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		t.Fatalf("create temp/test_file failed: %s", temp)
	}
	defer func() {
		_ = fileOpen.Close()
	}()

	fileNil := (*os.File)(nil)

	if !IsFileOpen(fileOpen) {
		t.Errorf("IsFileOpen(fileOpen) -> true: false")
	}

	if IsFileOpen(fileClose) {
		t.Errorf("IsFileOpen(fileClose) -> false: true")
	}

	if IsFileOpen(fileNil) {
		t.Errorf("IsFileOpen(fileNil) -> false: true")
	}
}
