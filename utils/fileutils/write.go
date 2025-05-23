// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fileutils

import (
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"io"
	"os"
	"strings"
)

func Write(filePath string, dat string) error {
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write([]byte(dat))
	return err
}

func WriteBytes(filePath string, dat []byte) error {
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write(dat)
	return err
}

func WriteEmpty(filePath string) error {
	// 尝试打开文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	return err
}

func AppendOnExistsFile(filePath string, dat string) error {
	// 尝试打开文件（若文件不存在则返回错误）
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write([]byte(dat))
	return err
}

// CheckFileByLine 检查文件每一行，若fn返回 true 则本函数返回 true，否则返回false。
// 传入 fn 的字符串不以 \n 结尾，空字符串也会传入
func CheckFileByLine(filePath string, fn func(string) bool) (bool, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	fileLine := strings.Split(cleanstringutils.GetString(string(dat)), "\n")
	for _, l := range fileLine {
		if fn(l) {
			return true, nil
		}
	}

	return false, nil
}

// ForEachLine 检查文件每一行。
// 传入 fn 的字符串不以 \n 结尾，空字符串也会传入
func ForEachLine(filePath string, fn func(string)) error {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileLine := strings.Split(cleanstringutils.GetString(string(dat)), "\n")
	for _, l := range fileLine {
		fn(l)
	}

	return nil
}

func Copy(dest string, src string) (int64, error) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = srcFile.Close()
	}()

	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = destFile.Close()
	}()

	// 使用 io.Copy 来复制内容
	nBytes, err := io.Copy(destFile, srcFile)
	if err != nil {
		return 0, err
	}

	return nBytes, nil
}
