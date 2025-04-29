// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package filesystemutils

import (
	"os"
	"path"
	"testing"
)

func TestIsExists(t *testing.T) {
	temp, err := os.MkdirTemp("", "test*")
	if err != nil {
		t.Fatalf("create temp directory failed: %s", temp)
	}
	defer func() {
		_ = os.RemoveAll(temp)
	}()

	testDir := path.Join(temp, "test_dir")
	err = os.Mkdir(testDir, 0700)
	if err != nil {
		t.Fatalf("create temp/test_dir failed: %s", temp)
	}

	testFile := path.Join(temp, "test_file")
	f, err := os.OpenFile(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		t.Fatalf("create temp/test_file failed: %s", temp)
	}
	_ = f.Close()

	testNotExists := path.Join(temp, "test_not_exists")

	// IsExists 测试

	if !IsExists(testFile) {
		t.Errorf("IsExists(%s) -> true: false", testFile)
	}

	if !IsExists(testDir) {
		t.Errorf("IsExists(%s) -> true: false", testDir)
	}

	if IsExists(testNotExists) {
		t.Errorf("IsExists(%s) -> false: true", testNotExists)
	}

	// IsFile 测试

	if !IsFile(testFile) {
		t.Errorf("IsFile(%s) -> true: false", testFile)
	}

	if IsFile(testDir) {
		t.Errorf("IsFile(%s) -> false: true", testDir)
	}

	if IsFile(testNotExists) {
		t.Errorf("IsFile(%s) -> false: true", testNotExists)
	}

	// IsExistsAndFile 测试

	if exists, file := IsExistsAndFile(testFile); !(exists && file) {
		t.Errorf("IsExistsAndFile(%s) -> true, true: %v, %v", testDir, exists, file)
	}

	if exists, file := IsExistsAndFile(testDir); !(exists && !file) {
		t.Errorf("IsExistsAndFile(%s) -> true, false: %v, %v", testDir, exists, file)
	}

	if exists, file := IsExistsAndFile(testNotExists); !(!exists && !file) {
		t.Errorf("IsExistsAndFile(%s) -> false, false: %v, %v", testDir, exists, file)
	}

	// IsDir 测试

	if IsDir(testFile) {
		t.Errorf("IsDir(%s) -> false: true", testFile)
	}

	if !IsDir(testDir) {
		t.Errorf("IsDir(%s) -> true: false", testDir)
	}

	if IsDir(testNotExists) {
		t.Errorf("IsDir(%s) -> false: true", testNotExists)
	}

	// IsExistsAndDir 测试

	if exists, dir := IsExistsAndDir(testFile); !(exists && !dir) {
		t.Errorf("IsExistsAndDir(%s) -> true, false: %v, %v", testDir, exists, dir)
	}

	if exists, dir := IsExistsAndDir(testDir); !(exists && dir) {
		t.Errorf("IsExistsAndDir(%s) -> true, true: %v, %v", testDir, exists, dir)
	}

	if exists, dir := IsExistsAndDir(testNotExists); !(!exists && !dir) {
		t.Errorf("IsExistsAndDir(%s) -> false, false: %v, %v", testDir, exists, dir)
	}
}
