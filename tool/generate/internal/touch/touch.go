// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package touch

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"os"
)

func TouchBaseFile() (err error) {
	genlog.InitGenLog("generate touch", os.Stdout)

	genlog.GenLog("touch base file")
	defer genlog.GenLog("touch base file finish")

	genlog.GenLogf("touch file %s\n", basefile.FileVersion)
	err = fileutils.Touch(basefile.FileVersion)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileLicense)
	err = fileutils.Touch(basefile.FileLicense)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileReport)
	err = fileutils.Touch(basefile.FileReport)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileBuildConfig)
	err = fileutils.Touch(basefile.FileBuildConfig)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileBuildDevConfig)
	err = fileutils.Touch(basefile.FileBuildDevConfig)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileBuildProdConfig)
	err = fileutils.Touch(basefile.FileBuildProdConfig)
	if err != nil {
		return err
	}

	genlog.GenLogf("touch file %s\n", basefile.FileBuildDateGob)
	err = fileutils.Touch(basefile.FileBuildDateGob)
	if err != nil {
		return err
	}

	return nil
}

func TouchReleaseFile() (err error) {
	genlog.GenLogf("touch file %s\n", basefile.FileReleaseInfoMD)
	err = fileutils.Touch(basefile.FileReleaseInfoMD)
	if err != nil {
		return err
	}

	return nil
}
