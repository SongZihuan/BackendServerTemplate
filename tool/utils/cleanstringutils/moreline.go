// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

import "strings"

func GetString(data string) (res string) {
	res = strings.Replace(data, "\r", "", -1)
	res = strings.Split(res, "\n")[0]
	res = strings.TrimSpace(res)
	return res
}
