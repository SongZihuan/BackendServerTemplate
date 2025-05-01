// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strconvutils

import "strconv"

func ParserInt(s string, base int, bitSize int) (i int64, err error) {
	res, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		return 0, err
	}

	return res, nil
}
