// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package randomutils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int, _charset string) string {
	if _charset == "" {
		return ""
	}

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = _charset[seededRand.Intn(len(_charset))]
	}
	return string(result)
}
