// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package version

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/utils/reutils"
	"strings"
	"sync"
)

var shortInitOnce sync.Once
var shortOnceErr error
var shortSemanticVersion = ""
var shortVersion = ""

func InitShortVersion(defVer string) error {
	shortInitOnce.Do(func() {
		shortOnceErr = initShortVersion(defVer)
	})

	return shortOnceErr
}

func initShortVersion(defVer string) error {
	genlog.GenLog("get short version info")
	defer genlog.GenLog("get short version info finish")

	var ver string

	err := git.InitGitData()
	if err != nil && !errors.Is(err, git.ErrWithoutGit) {
		return err
	}

	if err == nil {
		genlog.GenLog("try to get short version info from git")
		ver = getGitTagVersion()
		if ver != "" {
			shortSemanticVersion = ver
			shortVersion = "v" + shortSemanticVersion
			return nil
		} else {
			genlog.GenLogf("try to get short version info from git failed")
		}
	} else {
		genlog.GenLogf("try to get short version info from git failed: %s", err.Error())
	}

	genlog.GenLog("try to get short version info from file VERSION")
	ver = getDefaultVersion(defVer)
	if ver != "" {
		shortSemanticVersion = ver
		shortVersion = "v" + shortSemanticVersion
		return nil
	} else {
		genlog.GenLogf("try to get short version info from file VERSIOM failed")
	}

	genlog.GenLog("try to use pseudo short version")
	ver = getShortPseudoVersion()
	if ver != "" {
		shortSemanticVersion = ver
		shortVersion = "v" + shortSemanticVersion
		return nil
	} else {
		genlog.GenLogf("try to use pseudo short version failed")
	}

	genlog.GenLogf("get short version failed")

	return fmt.Errorf("get short version failed")
}

func getDefaultVersion(defVer string) string {
	defVer = strings.TrimPrefix(strings.ToLower(defVer), "v")
	if defVer == "" {
		return ""
	}

	defVer = strings.SplitN(defVer, "-", 2)[0]
	if defVer == "" || !reutils.IsSemanticVersion(defVer) {
		return ""
	}
	return defVer
}

func getGitTagVersion() (ver string) {
	if git.GetGitData().LastCommit == "" || git.GetGitData().LastTag == "" || git.GetGitData().LastTagCommit == "" {
		return ""
	}

	gitVer := strings.TrimPrefix(strings.ToLower(git.GetGitData().LastTag), "v")
	if ver = reutils.CheckAndGetShortSemanticVersion(gitVer); ver != "" {
		return gitVer
	}

	return ""
}

func getShortPseudoVersion() (randVer string) {
	return "0.0.0"
}

func WriteShortVersion(defVer string) error {
	genlog.GenLog("write short version data")
	defer genlog.GenLog("write short version data finish")

	err := InitShortVersion(defVer)
	if err != nil {
		genlog.GenLogf("get version failed: %s", err.Error())
		return nil
	}

	builder.SetShortVersion(shortVersion, shortSemanticVersion)

	return nil
}

func GetShortVersion(defVer string) string {
	err := InitShortVersion(defVer)
	if err != nil {
		genlog.GenLogf("get version failed: %s", err.Error())
		return ""
	}

	return shortVersion
}

func GetShortSemanticVersion(defVer string) string {
	err := InitShortVersion(defVer)
	if err != nil {
		genlog.GenLogf("get version failed: %s", err.Error())
		return ""
	}

	return shortSemanticVersion
}
