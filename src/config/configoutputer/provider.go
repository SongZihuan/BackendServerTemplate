// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configoutputer

import (
	"fmt"
	"strings"
)

const point = "."

var FileExtToProvider = map[string]NewProvider{
	point + "yaml": NewYamlProvider,
	point + "yml":  NewYamlProvider,
	point + "json": NewJsonProvider,
	point + "js":   NewJsonProvider,
}

func NewConfigOutputProvider(configPath string, opt *NewConfigOutputProviderOption) (ConfigOutputProvider, error) {
	for ext, provider := range FileExtToProvider {
		if strings.HasSuffix(configPath, ext) {
			return provider(opt), nil
		}
	}
	return nil, fmt.Errorf("config file type unknown")
}
