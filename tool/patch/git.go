// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/global"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/gitutils"
	"log"
	"strings"
	"sync"
)

const filePatchFile = "update.patch" + global.FileIgnoreExt

var onceGitInfo sync.Once
var toCommit string = ""
var fromCommit string = ""

var excludes = []string{
	// 配置文件
	"VERSION",
	"REPORT",
	"NAME",
	"LICENSE",
	"ENV_PREFIX",
	// 服务配置文件
	"SERVICE.yaml",
	// 文档
	"SECURITY.md",
	"README.md",
	"CONTRIBUTORS.md",
	"CONTRIBUTING.md",
	"CODE_OF_CONDUCT.md",
	"CHANGELOG_SPECIFICATION.md",
	"CHANGELOG.md",
	"dev-git-hooks/",
}

func InitGitData() (err error) {
	onceGitInfo.Do(func() {
		err = initGitData()
	})
	return err
}

func initGitData() (err error) {
	if !gitutils.HasGit() {
		log.Println("generate patch: `.git` not found, get git info skip")
		return nil
	}

	log.Println("generate patch: get git info")
	defer log.Println("generate patch: get git info finish")

	defer func() {
		if err != nil {
			toCommit = ""
			fromCommit = ""
		}
	}()

	tagList, err := gitutils.GetTagListWithFilter(func(s string) bool {
		return strings.HasPrefix(s, "v")
	})
	if err != nil {
		return err
	}
	log.Printf("generate patch: get git tag list length: %d\n", len(tagList))

	if len(tagList) == 0 {
		toCommit, err = gitutils.GetLastCommit()
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git to commit: %s\n", toCommit)

		fromCommit, err = gitutils.GetFirstCommit()
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git from commit: %s\n", fromCommit)
	} else if len(tagList) == 1 {
		toCommit, err = gitutils.GetTagCommit(tagList[0])
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git to commist (from tag '%s') : %s\n", tagList[0], toCommit)

		fromCommit, err = gitutils.GetFirstCommit()
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git from commit: %s\n", fromCommit)
	} else if len(tagList) >= 2 {
		toCommit, err = gitutils.GetTagCommit(tagList[0])
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git to commist (from tag '%s') : %s\n", tagList[0], toCommit)

		fromCommit, err = gitutils.GetTagCommit(tagList[1])
		if err != nil {
			return err
		}
		log.Printf("generate patch: get git from commist (from tag '%s') : %s\n", tagList[1], toCommit)
	} else {
		panic("unreachable")
	}

	return nil
}

func CreatePatch() error {
	err := InitGitData()
	if err != nil {
		return err
	}

	if !gitutils.HasGit() {
		return nil
	}

	log.Println("generate patch: create patch file")
	defer log.Println("generate patch: create patch file")

	if fromCommit == toCommit {
		log.Println("generate patch: only one commit, skip to create patch file")
		return fileutils.WriteEmpty(filePatchFile)
	}

	res, err := gitutils.GetPatch(fromCommit, toCommit, excludes...)
	if err != nil {
		return err
	}

	return fileutils.WriteBytes(filePatchFile, res)
}
