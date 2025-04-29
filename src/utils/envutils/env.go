// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envutils

import (
	"fmt"
	"os"
	"strings"
)

var EnvReplacer *strings.Replacer = nil

func init() {
	rules := make([]string, 0, (26+2)*2)
	rules = append(rules, ".", "_", "-", "_", " ", "")

	for _, i := range "abcdefghijklmnopqrstuvwxyz" {
		u := strings.ToUpper(string(i))
		rules = append(rules, string(i), u)
	}

	EnvReplacer = strings.NewReplacer(rules...)
}

func ToEnvName(input string) string {
	return EnvReplacer.Replace(input)
}

func GetSysEnv(name string) string {
	return os.Getenv(name)
}

func GetEnv(prefix string, name string) string {
	name = ToEnvName(name)

	if prefix != "" {
		return os.Getenv(fmt.Sprintf("%s_%s", prefix, name))
	}
	return os.Getenv(name)
}
