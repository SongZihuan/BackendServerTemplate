// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package randomutils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// GenerateRandomString generates a random string of the specified length
// containing lowercase letters ('a'-'z') and digits ('0'-'9').
func GenerateRandomString(length int, _charset string) string {
	if _charset == "" {
		_charset = charset
	}

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = _charset[seededRand.Intn(len(_charset))]
	}
	return string(result)
}
