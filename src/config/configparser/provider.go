// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"fmt"
	"strings"
)

func NewProvider(configPath string, opt *NewProviderOption) (ConfigParserProvider, error) {
	switch {
	case strings.HasSuffix(configPath, ".yaml") || strings.HasSuffix(configPath, ".yml"):
		return NewYamlProvider(opt), nil
	case strings.HasSuffix(configPath, ".yml"):
		return NewJsonProvider(opt), nil
	default:
		return nil, fmt.Errorf("config file type unknown")
	}
}
