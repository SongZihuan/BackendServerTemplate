// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package version

import (
	"errors"
	"fmt"
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/builder"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/builddate"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/random"
	"github.com/SongZihuan/BackendServerTemplate/utils/reutils"
	"strings"
	"sync"
)

var longInitOnce sync.Once
var longOnceErr error
var longSemanticVersion = ""
var longVersion = ""

func InitLongVersion() error {
	longInitOnce.Do(func() {
		longOnceErr = initLongVersion()
	})

	return longOnceErr
}

func initLongVersion() error {
	genlog.GenLog("get version info")
	defer genlog.GenLog("get version info finish")

	var ver string

	err := git.InitGitData()
	if err != nil && !errors.Is(err, git.ErrWithoutGit) {
		return err
	}

	if err == nil {
		genlog.GenLog("try to get version info from git")
		ver = getLongGitTagVersion()
		if ver != "" {
			longSemanticVersion = ver
			longVersion = "v" + longSemanticVersion
			return nil
		} else {
			genlog.GenLogf("try to get version info from git failed")
		}
	} else {
		genlog.GenLogf("try to get version info from git failed: %s", err.Error())
	}

	genlog.GenLog("try to get version info from file VERSION")
	ver = getLongDefaultVersion()
	if ver != "" {
		longSemanticVersion = ver
		longVersion = "v" + longSemanticVersion
		return nil
	} else {
		genlog.GenLogf("try to get version info from file VERSIOM failed")
	}

	genlog.GenLog("try to use pseudo long version")
	ver = getLongPseudoVersion()
	if ver != "" {
		longSemanticVersion = ver
		longVersion = "v" + longSemanticVersion
		return nil
	} else {
		genlog.GenLogf("try to use pseudo long version failed")
	}

	genlog.GenLogf("get long version failed")

	return fmt.Errorf("get long version failed")
}

func getLongDefaultVersion() (defVer string) {
	defVer = strings.TrimPrefix(strings.ToLower(resource.Version), "v")
	if defVer == "" || !reutils.IsSemanticVersion(defVer) {
		return ""
	}
	return defVer
}

func getLongGitTagVersion() (gitVer string) {
	gitVer = strings.TrimPrefix(strings.ToLower(git.GetGitData().LastTag), "v")
	if git.GetGitData().LastCommit != "" && (git.GetGitData().LastTagCommit == "" || gitVer == "") {
		// 存在当前提交，但提交没有任何
		return fmt.Sprintf("0.0.0+dev.%d.%s", builddate.GetBuildTime().Unix(), git.GetGitData().LastCommit)
	} else if git.GetGitData().LastCommit != "" && git.GetGitData().LastTagCommit != "" && gitVer != "" && reutils.IsSemanticVersion(gitVer) {
		if (git.GetGitData().LastCommit != git.GetGitData().LastTagCommit || strings.HasPrefix(gitVer, "0.")) && !strings.Contains(gitVer, "dev") {
			return gitVer + fmt.Sprintf("+dev.%d.%s", builddate.GetBuildTime().Unix(), git.GetGitData().LastCommit)
		}
		return gitVer
	} else {
		return ""
	}
}

func getLongPseudoVersion() (randVer string) {
	return fmt.Sprintf("0.0.0+dev.%d.%s", builddate.GetBuildTime().Unix(), random.GetPseudoCommitHash())
}

func WriteLongVersion() error {
	genlog.GenLog("write long version data")
	defer genlog.GenLog("write long version data finish")

	err := InitLongVersion()
	if err != nil {
		genlog.GenLogf("get version failed: %s", err.Error())
		return nil
	}

	builder.SetLongVersion(longVersion, longSemanticVersion)

	return nil
}

func GetLongVersion() string {
	return longVersion
}

func GetLongSemanticVersion() string {
	return longSemanticVersion
}
