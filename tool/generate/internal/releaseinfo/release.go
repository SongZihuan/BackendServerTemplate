// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package releaseinfo

import (
	_ "embed"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/changelog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/version"
	"github.com/SongZihuan/BackendServerTemplate/utils/reutils"
	"os"
	"strings"
	"text/template"
)

//go:embed release.md.tmpl
var releaseTemplateString string
var releaseTemplate *template.Template = template.New("Release")

//go:embed release.special.md.tmpl
var specialReleaseTemplateString string
var specialReleaseTemplate *template.Template = template.New("SpecialRelease")

type releaseTemplateData struct {
	Version       string
	GithubCompare string
	ChangeLog     string
}

type specialReleaseTemplateData struct {
	Version    string
	CommitHash string
}

func init() {
	var err error
	_, err = releaseTemplate.Parse(releaseTemplateString)
	if err != nil {
		panic(err)
	}

	_, err = specialReleaseTemplate.Parse(specialReleaseTemplateString)
	if err != nil {
		panic(err)
	}
}

func WriteReleaseData(defVer string) error {
	genlog.GenLog("write release info data data")
	defer genlog.GenLog("write write release info data finish")

	file, err := os.OpenFile(basefile.FileReleaseInfoMD, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	v := version.GetShortVersion(defVer)

	cl := changelog.GetLastChangLog(defVer)

	gcl := git.GetGitHubCompareMD()

	genlog.GenLogf("release info version: %s\n", v)
	if gcl != "" {
		genlog.GenLogf("release info github changelog: %s\n", gcl)
	} else {
		genlog.GenLog("release info github changelog: without")
	}
	genlog.GenLogf("release info changelog length=%d\n", len(cl))

	return releaseTemplate.Execute(file, releaseTemplateData{
		Version:       v,
		ChangeLog:     cl,
		GithubCompare: gcl,
	})
}

func WriteSpecialReleaseData(v string, force bool) error {
	genlog.GenLog("write special release info data data")
	defer genlog.GenLog("write write special release info data finish")

	file, err := os.OpenFile(basefile.FileReleaseInfoMD, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	if !reutils.IsSemanticVersion(v) {
		if !force {
			genlog.GenLogf("%s is not a semantic version", v)
			return fmt.Errorf("%s is not a semantic version", v)
		}
		genlog.GenLogf("warning: %s is not a semantic version", v)
	}

	commitHash := git.GetGitHubCommitMD()
	if commitHash == "" {
		commitHash = "*暂无*"
	}

	v = strings.TrimPrefix(strings.ToLower(v), "v")

	genlog.GenLogf("special release info version: %s\n", v)
	genlog.GenLogf("special release info commit hash: %s\n", commitHash)

	return specialReleaseTemplate.Execute(file, specialReleaseTemplateData{
		Version:    v,
		CommitHash: commitHash,
	})
}
