// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/global"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/gitutils"
	"os"
	"strings"
	"sync"
)

var onceGitInfo sync.Once
var lastCommit string = ""
var lastTag string = ""
var lastTagCommit string = ""
var secondToLastTag string = ""
var secondToLastTagCommit string = ""
var firstCommit string = ""

var onceIsGithub sync.Once
var isGithub bool = false

func InitGitData() (err error) {
	onceGitInfo.Do(func() {
		err = initGitData()
	})
	return err
}

func initGitData() (err error) {
	if !gitutils.HasGit() {
		genlog.GenLog("`.git` not found, get git info skip")
		return nil
	}

	genlog.GenLog("get git info")
	defer genlog.GenLog("get git info finish")

	defer func() {
		if err != nil {
			lastCommit = ""
			lastTag = ""
			lastTagCommit = ""
			secondToLastTag = ""
			secondToLastTagCommit = ""
			firstCommit = ""
		}
	}()

	lastCommit, err = gitutils.GetLastCommit()
	if err != nil {
		return err
	}
	genlog.GenLogf("get git last commit: %s\n", lastCommit)

	firstCommit, err = gitutils.GetFirstCommit()
	if err != nil {
		return err
	}
	genlog.GenLogf("get git first commit: %s\n", firstCommit)

	tagList, err := gitutils.GetTagListWithFilter(func(s string) bool {
		return strings.HasPrefix(s, "v")
	})
	if err != nil {
		return err
	}
	genlog.GenLogf("get git tag list length: %d\n", len(tagList))

	if len(tagList) > 0 {
		lastTag = tagList[0]
		genlog.GenLogf("get git last tag: %s\n", lastTag)

		lastTagCommit, err = gitutils.GetTagCommit(lastTag)
		if err != nil {
			return err
		}
		genlog.GenLogf("get git last tag commist: %s\n", lastTagCommit)
	} else {
		lastTag = ""
		lastTagCommit = ""
		genlog.GenLog("skip to get git last tag and last tag commit")
	}

	if len(tagList) > 1 {
		secondToLastTag = tagList[1]
		genlog.GenLogf("get git second to last tag: %s\n", secondToLastTag)

		secondToLastTagCommit, err = gitutils.GetTagCommit(secondToLastTag)
		if err != nil {
			return err
		}
		genlog.GenLogf("get git second to last tag commit: %s\n", secondToLastTagCommit)
	} else {
		secondToLastTag = ""
		secondToLastTagCommit = ""
		genlog.GenLog("skip to get git second to last tag and second to last tag commit")

	}

	return nil
}

func Version() string {
	_ = InitGitData()
	return lastTag
}

func WriteGitData() (err error) {
	genlog.GenLog("write git data")
	defer genlog.GenLog("write git data finish")

	_ = InitGitData()

	genlog.GenLogf("write %s to file %s\n", lastCommit, basefile.FileCommitDateTxt)
	err = fileutils.Write(basefile.FileCommitDateTxt, lastCommit)
	if err != nil {
		return err
	}

	genlog.GenLogf("write %s to file %s\n", lastTag, basefile.FileTagDataTxt)
	err = fileutils.Write(basefile.FileTagDataTxt, lastTag)
	if err != nil {
		return err
	}

	genlog.GenLogf("write %s to file %s\n", lastTagCommit, basefile.FileTagCommitData)
	err = fileutils.Write(basefile.FileTagCommitData, lastTagCommit)
	if err != nil {
		return err
	}

	return nil
}

func GetGitHubCompareMD() string {
	_ = InitGitData()

	compare, url := GetGitHubCompareURL()
	if compare == "" || url == "" {
		return ""
	}

	return fmt.Sprintf("**Git Changelog：** [%s](%s)", compare, url)
}

func GetGitHubCompareURL() (string, string) {
	_ = InitGitData()

	if !IsGitHub() || lastTag == "" || firstCommit == lastTagCommit {
		return "", ""
	}

	moduleName, err := mod.GetGoModuleName()
	if err != nil {
		return "", ""
	}

	if secondToLastTag != "" {
		compare := fmt.Sprintf("%s...%s", secondToLastTag, lastTag)
		return compare, fmt.Sprintf("https://%s/compare/%s", moduleName, compare)
	}

	return lastTag, fmt.Sprintf("https://%s/compare/%s...%s", moduleName, firstCommit, lastTag)
}

func IsGitHub() bool {
	onceIsGithub.Do(func() {
		moduleName, err := mod.GetGoModuleName()
		if err != nil {
			genlog.GenLog("the module is not on github, because module name not found")
			isGithub = false
			return
		}

		if strings.HasPrefix(moduleName, "github.com/") {
			genlog.GenLog("the module is on github")
			isGithub = true
		} else {
			genlog.GenLog("the module is not on github")
			isGithub = false
		}
	})
	return isGithub
}

func WriteGitIgnore() error {
	if !gitutils.HasGit() {
		genlog.GenLogf("`.git` not found, write %s skip\n", basefile.FileGitIgnore)
		return nil
	}

	genlog.GenLogf("write %s file\n", basefile.FileGitIgnore)
	defer genlog.GenLogf("write %s file finish\n", basefile.FileGitIgnore)

	s, err := os.Stat(basefile.FileGitIgnore)
	if err != nil {
		genlog.GenLogf("file %s not exists, create new one\n", basefile.FileGitIgnore)
		return newGitIgnore()
	}

	if s.IsDir() {
		genlog.GenLogf("%s is dir\n", basefile.FileGitIgnore)
		return fmt.Errorf("%s is dir", basefile.FileGitIgnore)
	}

	res, err := fileutils.CheckFileByLine(basefile.FileGitIgnore, func(s string) bool {
		if s == global.GitIgnoreExtFlag {
			return true
		}
		return false
	})
	if err != nil {
		return err
	} else if res {
		genlog.GenLogf("file %s check ok\n", basefile.FileGitIgnore)
		return nil
	}

	genlog.GenLogf("auto ignore '%s', write to file %s\n", global.GitIgnoreExtFlag, basefile.FileGitIgnore)
	return appendGitIgnore()
}

func newGitIgnore() error {
	return fileutils.Write(basefile.FileGitIgnore, fmt.Sprintf("# auto write by go generate\n%s\n", global.GitIgnoreExtFlag))
}

func appendGitIgnore() error {
	// 写入前添加\n，确保在新的一行
	return fileutils.AppendOnExistsFile(basefile.FileGitIgnore, fmt.Sprintf("\n# auto write by go generate\n%s\n", global.GitIgnoreExtFlag))
}
