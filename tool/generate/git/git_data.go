// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/mod"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/fileutils"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/gitutils"
	"log"
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
		log.Println("generate: `.git` not found, get git info skip")
		return nil
	}

	log.Println("generate: get git info")
	defer log.Println("generate: get git info finish")

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
	log.Printf("generate: get git last commit: %s\n", lastCommit)

	firstCommit, err = gitutils.GetFirstCommit()
	if err != nil {
		return err
	}
	log.Printf("generate: get git first commit: %s\n", firstCommit)

	tagList, err := gitutils.GetTagListWithFilter(func(s string) bool {
		return strings.HasPrefix(s, "v")
	})
	if err != nil {
		return err
	}

	log.Printf("generate: get git tag list length: %d\n", len(tagList))

	if len(tagList) > 0 {
		lastTag = tagList[0]
		log.Printf("generate: get git last tag: %s\n", lastTag)

		lastTagCommit, err = gitutils.GetTagCommit(lastTag)
		if err != nil {
			return err
		}
		log.Printf("generate: get git last tag commist: %s\n", lastTagCommit)
	} else {
		lastTag = ""
		lastTagCommit = ""
		log.Println("generate: skip to get git last tag and last tag commit")
	}

	if len(tagList) > 1 {
		secondToLastTag = tagList[1]
		log.Printf("generate: get git second to last tag: %s\n", secondToLastTag)

		secondToLastTagCommit, err = gitutils.GetTagCommit(secondToLastTag)
		if err != nil {
			return err
		}
		log.Printf("generate: get git second to last tag commit: %s\n", secondToLastTagCommit)
	} else {
		secondToLastTag = ""
		secondToLastTagCommit = ""
		log.Println("generate: skip to get git second to last tag and second to last tag commit")

	}

	return nil
}

func Version() string {
	_ = InitGitData()
	return lastTag
}

func WriteGitData() (err error) {
	log.Println("generate: write git data")
	defer log.Println("generate: write git data finish")

	_ = InitGitData()

	log.Printf("generate: write %s to file %s\n", lastCommit, basefile.FileCommitDateTxt)
	err = fileutils.Write(basefile.FileCommitDateTxt, lastCommit)
	if err != nil {
		return err
	}

	log.Printf("generate: write %s to file %s\n", lastTag, basefile.FileTagDataTxt)
	err = fileutils.Write(basefile.FileTagDataTxt, lastTag)
	if err != nil {
		return err
	}

	log.Printf("generate: write %s to file %s\n", lastTagCommit, basefile.FileTagCommitData)
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
			log.Println("generate: the module is not on github, because module name not found")
			isGithub = false
			return
		}

		if strings.HasPrefix(moduleName, "github.com/") {
			log.Println("generate: the module is on github")
			isGithub = true
		} else {
			log.Println("generate: the module is not on github")
			isGithub = false
		}
	})
	return isGithub
}

func WriteGitIgnore() error {
	if !gitutils.HasGit() {
		log.Printf("generate: `.git` not found, write %s skip\n", basefile.FileGitIgnore)
		return nil
	}

	log.Printf("generate: write %s file\n", basefile.FileGitIgnore)
	defer log.Printf("generate: write %s file finish\n", basefile.FileGitIgnore)

	s, err := os.Stat(basefile.FileGitIgnore)
	if err != nil {
		log.Printf("generaate: file %s not exists, create new one\n", basefile.FileGitIgnore)
		return newGitIgnore()
	}

	if s.IsDir() {
		log.Printf("generaate: %s is dir\n", basefile.FileGitIgnore)
		return fmt.Errorf("%s is dir", basefile.FileGitIgnore)
	}

	res, err := fileutils.CheckFileByLine(basefile.FileGitIgnore, func(s string) bool {
		if s == basefile.GitIgnoreExtFlag {
			return true
		}
		return false
	})
	if err != nil {
		return err
	} else if res {
		log.Printf("generaate: file %s check ok\n", basefile.FileGitIgnore)
		return nil
	}

	log.Printf("generaate: auto ignore '%s', write to file %s\n", basefile.GitIgnoreExtFlag, basefile.FileGitIgnore)
	return appendGitIgnore()
}

func newGitIgnore() error {
	return fileutils.Write(basefile.FileGitIgnore, fmt.Sprintf("# auto write by go generate\n%s\n", basefile.GitIgnoreExtFlag))
}

func appendGitIgnore() error {
	// 写入前添加\n，确保在新的一行
	return fileutils.AppendOnExistsFile(basefile.FileGitIgnore, fmt.Sprintf("\n# auto write by go generate\n%s\n", basefile.GitIgnoreExtFlag))
}
