// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strconvutils

import "testing"

func TestParserInt(t *testing.T) {
	if res, err := ParserInt("100", 10, 64); err != nil {
		t.Errorf("ParserInt(100) error: %s", err.Error())
	} else if res != 100 {
		t.Errorf("ParserInt(100) -> 100: %v", res)
	}

	if res, err := ParserInt("abc", 10, 64); err == nil {
		t.Errorf("ParserInt(abc) error: err is nil")
	} else if res != 0 {
		t.Errorf("ParserInt(abc) -> 0: %v", res)
	}
}
