// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package random

import (
	"github.com/SongZihuan/BackendServerTemplate/utils/randomutils"
)

var pseudoCommitHash = ""

const commitHashCharset = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	pseudoCommitHash = randomutils.GenerateRandomString(40, commitHashCharset)
}

func GetPseudoCommitHash() string {
	return pseudoCommitHash
}
