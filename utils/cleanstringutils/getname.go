// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

import "unicode"

func GetName(input string) string {
	var result []rune

	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '-' || char == '_' || char == '.' {
			result = append(result, char)
		}
	}

	return string(result)
}
