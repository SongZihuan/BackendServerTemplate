// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package releaseinfo

import (
	_ "embed"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/changelog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/genlog"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/git"
	"os"
	"text/template"
)

//go:embed release_data.md.tmpl
var releaseTemplateString string
var releaseTemplate *template.Template = template.New("Release")

type templateData struct {
	Version       string
	GithubCompare string
	ChangeLog     string
}

func init() {
	var err error
	_, err = releaseTemplate.Parse(releaseTemplateString)
	if err != nil {
		panic(err)
	}
}

func WriteReleaseData() error {
	genlog.GenLog("write release info data data")
	defer genlog.GenLog("write write release info data finish")

	file, err := os.OpenFile(basefile.FileReleaseInfoMD, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	v := git.Version()

	cl := changelog.GetLastChangLog()

	gcl := git.GetGitHubCompareMD()

	genlog.GenLogf("release info version: %s\n", v)
	if gcl != "" {
		genlog.GenLogf("release info github changelog: %s\n", gcl)
	} else {
		genlog.GenLog("release info github changelog: without")
	}
	genlog.GenLogf("release info changelog length=%d\n", len(cl))

	return releaseTemplate.Execute(file, templateData{
		Version:       git.Version(),
		ChangeLog:     changelog.GetLastChangLog(),
		GithubCompare: git.GetGitHubCompareMD(),
	})
}
