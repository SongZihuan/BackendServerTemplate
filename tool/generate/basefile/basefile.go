// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basefile

import "github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"

const (
	FileVersion   = "./VERSION"
	FileLicense   = "./LICENSE"
	FileReport    = "./REPORT"
	FileName      = "./NAME"
	FileEnvPrefix = "./ENV_PREFIX"

	FileBuildDateTxt  = "./build_date.txt"
	FileCommitDateTxt = "./commit_data.txt"
	FileTagDataTxt    = "./tag_data.txt"
	FileTagCommitData = "./tag_commit_data.txt"
	FileRandomData    = "./random_data.txt"

	FileReleaseInfoMD = "./release_info.md"
)

func CreateBaseFile() (err error) {
	err = fileutils.Touch(FileVersion)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileLicense)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileReport)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileName)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileEnvPrefix)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileBuildDateTxt)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileCommitDateTxt)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileTagDataTxt)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileTagCommitData)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileRandomData)
	if err != nil {
		return err
	}

	err = fileutils.Touch(FileReleaseInfoMD)
	if err != nil {
		return err
	}

	return nil
}
