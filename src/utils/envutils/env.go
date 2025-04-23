// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envutils

import (
	"fmt"
	resource "github.com/SongZihuan/BackendServerTemplate"
	"os"
	"strings"
)

var EnvPrefix = resource.EnvPrefix

func init() {
	if EnvPrefix == "" {
		return
	}

	newEnvPrefix := StringToEnvName(EnvPrefix)
	if EnvPrefix != newEnvPrefix {
		panic(fmt.Errorf("bad %s; good %s", EnvPrefix, newEnvPrefix))
	} else if strings.HasSuffix(EnvPrefix, "_") {
		panic("EnvPrefix End With '_'")
	}
}

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
	rules = append(rules, ".", "_", "-", "_", " ", "")

	for _, i := range "abcdefghijklmnopqrstuvwxyz" {
		u := strings.ToUpper(string(i))
		rules = append(rules, string(i), u)
	}

	return strings.NewReplacer(rules...)
}

func GetSysEnv(name string) string {
	return os.Getenv(name)
}

func GetEnv(name string) string {
	if resource.EnvPrefix != "" {
		return os.Getenv(fmt.Sprintf("%s_%s", resource.EnvPrefix, name))
	}
	return os.Getenv(name)
}
