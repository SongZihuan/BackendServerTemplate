// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envutils

import (
	"strings"
)

func StringToEnvName(input string) string {
	// Replace '.' and '-' with '_'
	replaced := strings.NewReplacer(".", "_", "-", "_").Replace(input)

	// Convert to uppercase
	upper := strings.ToUpper(replaced)

	// Remove all other symbols
	cleaned := strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_' {
			return r
		}
		return -1
	}, upper)

	return cleaned
}

func GetEnvReplaced() *strings.Replacer {
	rules := make([]string, 0, (26+2)*2)
	rules = append(rules, ".", "_", "-", "_")

	for _, i := range "abcdefghijklmnopqrstuvwxyz" {
		u := strings.ToUpper(string(i))
		rules = append(rules, string(i), u)
	}

	return strings.NewReplacer(rules...)
}
