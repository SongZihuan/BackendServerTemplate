// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resource

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

const (
	FromNo      = "no"
	FromInstall = "install"
	FromConfig  = "config"
)

const (
	Args1Install    = "install"
	Args1Uninstall1 = "remove"
	Args1Uninstall2 = "uninstall"
	Args1Start      = "start"
	Args1Stop       = "stop"
	Args1Restart    = "restart"
)

type ServiceConfigType struct {
	Name         string            `yaml:"name"`
	DisplayName  string            `yaml:"display-name"`
	Describe     string            `yaml:"describe"`
	ArgumentFrom string            `yaml:"argument-from"`
	ArgumentList []string          `yaml:"argument-list"`
	EnvFrom      string            `yaml:"env-from"`
	EnvGetList   []string          `yaml:"env-get-list"`
	EnvSetList   map[string]string `yaml:"env-set-list"`
}

//go:embed SERVICE.yaml
var serviceConfig []byte
var ServiceConfig ServiceConfigType

var nameRegex = regexp.MustCompilePOSIX(`^[a-zA-Z0-9]+$`)

func initServiceConfig() {
	err := yaml.Unmarshal(serviceConfig, &ServiceConfig)
	if err != nil {
		panic(err)
	}

	if ServiceConfig.Name == "" || !nameRegex.MatchString(ServiceConfig.Name) {
		panic("service name is invalid")
	}

	if ServiceConfig.DisplayName == "" {
		ServiceConfig.DisplayName = ServiceConfig.Name
	}

	ServiceConfig.Describe = utilsClenFileData(ServiceConfig.Describe)

	switch ServiceConfig.ArgumentFrom {
	case FromInstall:
		if len(os.Args) > 2 && os.Args[1] == Args1Install {
			ServiceConfig.ArgumentFrom = FromConfig
			ServiceConfig.ArgumentList = os.Args[2:]
		} else {
			ServiceConfig.ArgumentFrom = FromNo
			ServiceConfig.ArgumentList = nil
		}
	case FromConfig:
		if len(ServiceConfig.ArgumentList) == 0 {
			ServiceConfig.ArgumentFrom = FromNo
			ServiceConfig.ArgumentList = nil
		}
	default:
		ServiceConfig.ArgumentFrom = FromNo
		ServiceConfig.ArgumentList = nil
	}

	switch ServiceConfig.EnvFrom {
	case FromInstall:
		if len(ServiceConfig.EnvGetList) == 0 {
			ServiceConfig.EnvFrom = FromNo
			ServiceConfig.EnvGetList = nil
			ServiceConfig.EnvSetList = nil
			break
		}

		ServiceConfig.EnvSetList = make(map[string]string, len(ServiceConfig.EnvGetList))
		for _, e := range ServiceConfig.EnvGetList {
			ServiceConfig.EnvSetList[e] = os.Getenv(e)
		}
	case FromConfig:
		ServiceConfig.EnvGetList = nil
		if len(ServiceConfig.EnvSetList) == 0 {
			ServiceConfig.EnvFrom = FromNo
			ServiceConfig.EnvSetList = nil
		}
	default:
		ServiceConfig.EnvFrom = FromNo
		ServiceConfig.EnvGetList = nil
		ServiceConfig.EnvSetList = nil
	}
}
