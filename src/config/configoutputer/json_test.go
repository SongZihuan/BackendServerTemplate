// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configoutputer

import "testing"

func TestJson(t *testing.T) {
	var a ConfigOutputProvider
	a = &JsonProvider{}
	_ = a
}

func TestNewJsonFunc(t *testing.T) {
	var a NewProvider
	a = NewJsonProvider
	_ = a
}
