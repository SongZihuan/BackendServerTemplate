// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/executils"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
)

func winMTManifest(exe string, manifest string) error {
	_, err := executils.Run(mtptogram, "-manifest", manifest, fmt.Sprintf("-outputresource:%s;1", exe))
	if err != nil {
		return err
	}
	return nil
}

func winCopyManifest(exe string, manifest string) error {
	_, err := fileutils.Copy(fmt.Sprintf("%s.manifest", exe), manifest) // exe 可以是路径，也可以是文件
	if err != nil {
		return err
	}
	return nil
}
