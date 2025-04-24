// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/builddate"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/random"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/releaseinfo"
	"log"
	"os"
)

func main() {
	os.Exit(command())
}

func command() int {
	var err error

	log.Println("generate start to run")

	_, err = mod.GetGoModuleName() // 提前一步帕胺的
	if err != nil {
		return ReturnError(err)
	}

	err = basefile.TouchBaseFile()
	if err != nil {
		return ReturnError(err)
	}

	err = git.InitGitData()
	if err != nil {
		return ReturnError(err)
	}

	err = builddate.WriteBuildDateData()
	if err != nil {
		return ReturnError(err)
	}

	err = git.WriteGitData()
	if err != nil {
		return ReturnError(err)
	}

	err = git.WriteGitIgnore()
	if err != nil {
		return ReturnError(err)
	}

	err = random.WriteRandomData()
	if err != nil {
		return ReturnError(err)
	}

	err = releaseinfo.WriteReleaseData()
	if err != nil {
		return ReturnError(err)
	}

	return ReturnSuccess()
}
