// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/gitutils"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/modutils"
	"strings"
	"sync"
)

var once sync.Once
var lastCommit string = ""
var lastTag string = ""
var lastTagCommit string = ""
var secondToLastTag string = ""
var secondToLastTagCommit string = ""
var firstCommit string = ""

func InitGitData() (err error) {
	once.Do(func() {
		err = initGitData()
	})
	return err
}

func initGitData() (err error) {
	if !gitutils.HasGit() {
		return nil
	}

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

	firstCommit, err = gitutils.GetFirstCommit()
	if err != nil {
		return err
	}

	tagList, err := gitutils.GetTagListWithFilter(func(s string) bool {
		return strings.HasPrefix(s, "v")
	})
	if err == nil {
		return err
	}

	if len(tagList) > 0 {
		lastTag = tagList[0]
		lastTagCommit, err = gitutils.GetTagCommit(lastTag)
		if err != nil {
			return err
		}
	}

	if len(tagList) > 1 {
		secondToLastTag = tagList[1]
		secondToLastTagCommit, err = gitutils.GetTagCommit(secondToLastTag)
		if err != nil {
			return err
		}
	} else {
		secondToLastTag = ""
		secondToLastTagCommit = ""
	}

	return nil
}

func Version() string {
	_ = InitGitData()
	return lastTag
}

func WriteGitData() (err error) {
	_ = InitGitData()

	err = fileutils.Write(basefile.FileCommitDateTxt, lastCommit)
	if err != nil {
		return err
	}

	err = fileutils.Write(basefile.FileTagDataTxt, lastTag)
	if err != nil {
		return err
	}

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

	return fmt.Sprintf("**Git Changelog：[%s](%s)", compare, url)
}

func GetGitHubCompareURL() (string, string) {
	_ = InitGitData()

	if !modutils.IsGitHub || lastTag == "" || firstCommit == lastTagCommit {
		return "", ""
	}

	if secondToLastTag != "" {
		compare := fmt.Sprintf("%s...%s", secondToLastTag, lastTag)
		return compare, fmt.Sprintf("https://%s/compare/%s", modutils.ModPath, compare)
	}

	return lastTag, fmt.Sprintf("https://%s/compare/%s...%s", modutils.ModPath, firstCommit, lastTag)
}
