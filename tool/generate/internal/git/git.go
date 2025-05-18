// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/mod"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/gitutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/reutils"
	"os"
	"strings"
	"sync"
)

var ErrWithoutGit = fmt.Errorf("without git")

type GitData struct {
	LastCommit            string
	FirstCommit           string
	LastTag               string
	LastTagCommit         string
	SecondToLastTag       string
	SecondToLastTagCommit string
}

var onceGitInfo sync.Once
var onceGitInfoErr error
var data GitData

var onceIsGithub sync.Once
var isGithub bool = false

func InitGitData() error {
	onceGitInfo.Do(func() {
		onceGitInfoErr = initGitData()
	})
	return onceGitInfoErr
}

func initGitData() (err error) {
	if !gitutils.HasGit() {
		genlog.GenLog("`.git` not found, get git info skip")
		return ErrWithoutGit
	}

	genlog.GenLog("get git info")
	defer genlog.GenLog("get git info finish")

	defer func() {
		if err != nil {
			data = GitData{}
		}
	}()

	data.LastCommit, err = gitutils.GetLastCommit()
	if err != nil {
		return err
	}
	genlog.GenLogf("get git last commit: %s\n", data.LastCommit)

	data.FirstCommit, err = gitutils.GetFirstCommit()
	if err != nil {
		return err
	}
	genlog.GenLogf("get git first commit: %s\n", data.FirstCommit)

	tagList, err := gitutils.GetTagListWithFilter(func(s string) bool {
		if !strings.HasPrefix(s, "v") {
			return false
		}

		if v := strings.TrimPrefix(s, "v"); !reutils.IsSemanticVersion(v) {
			return false
		}

		return true
	})
	if err != nil {
		return err
	}
	genlog.GenLogf("get git tag list length: %d\n", len(tagList))

	if len(tagList) > 0 {
		data.LastTag = tagList[0]
		genlog.GenLogf("get git last tag: %s\n", data.LastTag)

		data.LastTagCommit, err = gitutils.GetTagCommit(data.LastTag)
		if err != nil {
			return err
		}
		genlog.GenLogf("get git last tag commist: %s\n", data.LastTagCommit)
	} else {
		data.LastTag = ""
		data.LastTagCommit = ""
		genlog.GenLog("skip to get git last tag and last tag commit")
	}

	if len(tagList) > 1 {
		data.SecondToLastTag = tagList[1]
		genlog.GenLogf("get git second to last tag: %s\n", data.SecondToLastTag)

		data.SecondToLastTagCommit, err = gitutils.GetTagCommit(data.SecondToLastTag)
		if err != nil {
			return err
		}
		genlog.GenLogf("get git second to last tag commit: %s\n", data.SecondToLastTagCommit)
	} else {
		data.SecondToLastTag = ""
		data.SecondToLastTagCommit = ""
		genlog.GenLog("skip to get git second to last tag and second to last tag commit")

	}

	return nil
}

func GetGitData() GitData {
	err := InitGitData()
	if err != nil {
		return GitData{}
	}

	return data
}

func WriteGitData() (err error) {
	genlog.GenLog("write git data")
	defer genlog.GenLog("write git data finish")

	err = InitGitData()
	if errors.Is(err, ErrWithoutGit) {
		genlog.GenLog("without git, skip write git data")
		return nil
	} else if err != nil {
		genlog.GenLogf("init git info failed: %s", err.Error())
		return nil
	}

	builder.SetCommitHash(data.LastCommit)

	return nil
}

func GetGitHubCompareMD() string {
	err := InitGitData()
	if err != nil {
		return ""
	}

	if !IsGitHub() || data.LastTag == "" || data.FirstCommit == data.LastTagCommit {
		return ""
	}

	compare, url := getGitHubCompareURL()
	if compare == "" || url == "" {
		return ""
	}

	return fmt.Sprintf("**Git Changelog：** [%s](%s)", compare, url)
}

func GetGitHubCompareURL() (string, string) {
	err := InitGitData()
	if err != nil {
		return "", ""
	}

	if !IsGitHub() || data.LastTag == "" || data.FirstCommit == data.LastTagCommit {
		return "", ""
	}

	return getGitHubCompareURL()
}

func getGitHubCompareURL() (string, string) {
	moduleName, err := mod.GetGoModuleName()
	if err != nil {
		return "", ""
	}

	if data.SecondToLastTag != "" {
		compare := fmt.Sprintf("%s...%s", data.SecondToLastTag, data.LastTag)
		return compare, fmt.Sprintf("https://%s/compare/%s", moduleName, compare)
	}

	return data.LastTag, fmt.Sprintf("https://%s/compare/%s...%s", moduleName, data.FirstCommit, data.LastTag)
}

func GetGitHubCommitMD() string {
	err := InitGitData()
	if err != nil {
		return ""
	}

	if !IsGitHub() || data.LastCommit == "" {
		return ""
	}

	if len(data.LastCommit) >= 8 {
		data.LastCommit = data.LastCommit[:8]
	}

	url := getGitHubCommitURL()
	if url != "" {
		return ""
	}

	return fmt.Sprintf("[%s](%s)", data.LastCommit, url)
}

func GetGitHubCommitURL() string {
	err := InitGitData()
	if err != nil {
		return ""
	}

	if !IsGitHub() || data.LastCommit == "" {
		return ""
	}

	return getGitHubCommitURL()
}

func getGitHubCommitURL() string {
	moduleName, err := mod.GetGoModuleName()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("https://%s/commit/%s", moduleName, data.LastCommit)
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

	var ignoreFlags = make(map[string]bool, len(basefile.GitIgnoreList))

	err = fileutils.ForEachLine(basefile.FileGitIgnore, func(s string) {
		for _, f := range basefile.GitIgnoreList {
			if f == s {
				ignoreFlags[s] = true
				break
			}
		}
	})
	if err != nil {
		return err
	} else if len(ignoreFlags) == len(basefile.GitIgnoreList) {
		genlog.GenLogf("file %s check ok\n", basefile.FileGitIgnore)
		return nil
	}

	genlog.GenLogf("auto add ignore file to %s\n", basefile.FileGitIgnore)
	return appendGitIgnore(ignoreFlags)
}

func newGitIgnore() error {
	var res strings.Builder

	moduleName, err := mod.GetGoModuleName()
	if err != nil {
		return err
	}

	res.WriteString(fmt.Sprintf("# Automatically generated by go project %s\n", moduleName))
	for _, f := range basefile.GitIgnoreList {
		res.WriteString(fmt.Sprintf("%s\n", f))
	}

	return fileutils.Write(basefile.FileGitIgnore, res.String())
}

func appendGitIgnore(existsFlags map[string]bool) error {
	var res strings.Builder

	moduleName, err := mod.GetGoModuleName()
	if err != nil {
		return err
	}

	// 写入前添加\n，确保在新的一行
	res.WriteString(fmt.Sprintf("\n# Automatically generated by go project %s\n", moduleName))
	for _, f := range basefile.GitIgnoreList {
		yes, ok := existsFlags[f]

		fmt.Println("TAG A", f, yes, ok)

		if ok && yes {
			continue
		}

		res.WriteString(fmt.Sprintf("%s\n", f))
	}

	return fileutils.AppendOnExistsFile(basefile.FileGitIgnore, res.String())
}
