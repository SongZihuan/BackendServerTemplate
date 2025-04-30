// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configerror

import "testing"

func ErrorTest(t *testing.T) { // 测试函数，确保 Error 符合 error
	var a Error
	var b error

	b = a
	_ = b
}
