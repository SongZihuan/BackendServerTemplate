// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basefile

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/global"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"log"
)

const (
	FileVersion    = "./VERSION"
	FileLicense    = "./LICENSE"
	FileReport     = "./REPORT"
	FileName       = "./NAME"
	FileEnvPrefix  = "./ENV_PREFIX"
	FileSystemYaml = "./SERVICE.yaml"

	FileBuildDateTxt  = "./build_date.dat" + global.FileIgnoreExt
	FileCommitDateTxt = "./commit_data.dat" + global.FileIgnoreExt
	FileTagDataTxt    = "./tag_data.dat" + global.FileIgnoreExt
	FileTagCommitData = "./tag_commit_data.dat" + global.FileIgnoreExt
	FileRandomData    = "./random_data.dat" + global.FileIgnoreExt

	FileReleaseInfoMD = "./release_info.md" + global.FileIgnoreExt

	FileGitIgnore = "./.gitignore"
)

func TouchBaseFile() (err error) {
	log.Println("generate: touch base file")
	defer log.Println("generate: touch base file finish")

	log.Printf("generate: touch file %s\n", FileVersion)
	err = fileutils.Touch(FileVersion)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileLicense)
	err = fileutils.Touch(FileLicense)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileReport)
	err = fileutils.Touch(FileReport)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileName)
	err = fileutils.Touch(FileName)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileEnvPrefix)
	err = fileutils.Touch(FileEnvPrefix)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileSystemYaml)
	err = fileutils.Touch(FileSystemYaml)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileBuildDateTxt)
	err = fileutils.Touch(FileBuildDateTxt)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileCommitDateTxt)
	err = fileutils.Touch(FileCommitDateTxt)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileTagDataTxt)
	err = fileutils.Touch(FileTagDataTxt)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileTagCommitData)
	err = fileutils.Touch(FileTagCommitData)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileRandomData)
	err = fileutils.Touch(FileRandomData)
	if err != nil {
		return err
	}

	log.Printf("generate: touch file %s\n", FileReleaseInfoMD)
	err = fileutils.Touch(FileReleaseInfoMD)
	if err != nil {
		return err
	}

	return nil
}
