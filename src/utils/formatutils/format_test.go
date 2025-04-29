// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package formatutils

import (
	_ "embed"
	"github.com/SongZihuan/BackendServerTemplate/tool/utils/cleanstringutils"
	"strings"
	"testing"
)

//go:embed test_file_1.txt
var testFileInput string

//go:embed test_file_2.txt
var testFileOutput string

func init() {
	testFileInput = cleanstringutils.GetString(testFileInput)
	testFileOutput = strings.TrimRight(strings.Replace(testFileOutput, "\r", "\n", -1), "\n")
}

func TestFormat(t *testing.T) {
	res := FormatTextToWidthAndPrefix(testFileInput, 10, 80)
	if res != testFileOutput {
		t.Errorf("format string error: \n===START EXPECTED===\n%s\n===END===\n===START ACTUAL===\n%s\n===END===\n", testFileOutput, res)
	}
}
