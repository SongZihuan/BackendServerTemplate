// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func createTarGZip(dest string, src string) error {
	// 创建目标.tar.gz文件
	file, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	// 创建gzip写入器
	gzipWriter := gzip.NewWriter(file)
	defer func() {
		_ = gzipWriter.Close()
	}()

	// 创建tar写入器
	tarWriter := tar.NewWriter(gzipWriter)
	defer func() {
		_ = tarWriter.Close()
	}()

	// 遍历源目录
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") { // 跳过此文件
			return nil
		}

		if strings.HasPrefix(info.Name(), "_") { // 跳过此文件
			return nil
		}

		if strings.HasSuffix(info.Name(), "_") { // 跳过此文件
			return nil
		}

		// 创建一个头信息
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// 确保我们使用的是相对于sourceDir的相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 清理敏感信息
		header = headerClean(header)

		// 写入头信息
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// 如果是目录，则跳过写入内容
		if info.IsDir() {
			return nil
		}

		// 如果是文件，则写入内容
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()

		_, err = io.Copy(tarWriter, f)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func headerClean(header *tar.Header) *tar.Header {
	header.Mode = 0644
	header.Uid = 0
	header.Gid = 0
	header.Uname = ""
	header.Gname = ""

	header.ModTime = time.Now()
	header.ChangeTime = time.Time{}
	header.AccessTime = time.Time{}

	header.Xattrs = nil // 已废除
	header.PAXRecords = nil

	header.Format = tar.FormatUSTAR // 尽可能的通用

	return header
}
