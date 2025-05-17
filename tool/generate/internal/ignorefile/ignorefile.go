// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ignorefile

import "github.com/SongZihuan/BackendServerTemplate/tool/generate/internal/basefile"

const (
	IgnoreExt        = ".ignore" // 一般的被忽略文件
	PatchExt         = ".patch"  // 补丁文件（默认情况下忽略）
	BuildMarkdownExt = ".ot.md"  // 自动生成的一次性Markdown文件
	BuildDataExt     = ".otd"    // 自动生成的一次性编译信息文件
)

const GitIgnoreStar = "*"

const (
	GitIgnoreIgnoreExt        = GitIgnoreStar + IgnoreExt
	GitIgnorePatchExt         = GitIgnoreStar + PatchExt
	GitIgnoreBuildMarkdownExt = GitIgnoreStar + BuildMarkdownExt
	GitIgnoreBuildDataExt     = GitIgnoreStar + BuildDataExt
	GitIgnoreBuildConfigFile  = basefile.FileBuildConfig
)

var GitIgnoreList = []string{
	GitIgnoreIgnoreExt,
	GitIgnorePatchExt,
	GitIgnoreBuildMarkdownExt,
	GitIgnoreBuildDataExt,
}
