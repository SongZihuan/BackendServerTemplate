// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitutils

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/executils"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"strings"
	"sync"
)

var hasGitOnce sync.Once
var hasGit = false

func HasGit() bool {
	hasGitOnce.Do(func() {
		hasGit = filesystemutils.IsDir("./.git")
	})
	return hasGit
}

func GetLastCommit() (string, error) {
	return executils.RunOnline("git", "rev-parse", "HEAD")
}

func GetTagListWithFilter(filter func(string) bool) ([]string, error) {
	ret, err := executils.Run("git", "for-each-ref", "--sort=-creatordate", "--format", "%(refname:short)", "refs/tags/")
	if err != nil {
		return nil, err
	}

	ret = cleanstringutils.GetString(ret)

	tagListSrc := strings.Split(ret, "\n")

	tagList := make([]string, 0, len(tagListSrc))

	for _, tag := range tagListSrc {
		tag = strings.TrimSpace(tag)
		if tag == "" || (filter != nil && !filter(tag)) {
			continue
		}
		tagList = append(tagList, tag)
	}

	return tagList, nil
}

func GetTagList() ([]string, error) {
	return GetTagListWithFilter(nil)
}

func GetTagCommit(tag string) (string, error) {
	return executils.RunOnline("git", "rev-list", "-n", "1", tag)
}

func GetFirstCommit() (string, error) {
	return executils.RunOnline("git", "rev-list", "--max-parents=0", "HEAD")
}

func GetPatch(from string, to string, excludes ...string) ([]byte, error) {
	args := make([]string, 0, len(excludes)+3)
	args = append(args, "diff", from, to, ".")
	for _, e := range excludes {
		args = append(args, fmt.Sprintf(":(exclude)%s", e))
	}

	return executils.RunBytes("git", args...)
}
