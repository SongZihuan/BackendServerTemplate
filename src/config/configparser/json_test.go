// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import "testing"

func TestJson(t *testing.T) {
	var a ConfigParserProvider
	a = &JsonProvider{}
	_ = a
}
