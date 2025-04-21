// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/sliceutils"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

const (
	FromNo      = "no"
	FromInstall = "install"
	FromConfig  = "config"
)

type ServiceConfigType struct {
	Name         string            `yaml:"name,omitempty"`
	DisplayName  string            `yaml:"display-name,omitempty"`
	Describe     string            `yaml:"describe,omitempty"`
	ArgumentFrom string            `yaml:"argument-from,omitempty"`
	ArgumentList []string          `yaml:"argument-list,omitempty"`
	EnvFrom      string            `yaml:"env-from,omitempty"`
	EnvGetList   []string          `yaml:"env-get-list,omitempty"`
	EnvSetList   map[string]string `yaml:"env-set-list,omitempty"`
}

var serviceConfigYamlData = global.ServiceConfigYamlData
var serviceConfig ServiceConfigType

var nameRegex = regexp.MustCompilePOSIX(`^[a-zA-Z0-9]+$`)

func initInstallServiceConfig(args []string) error {
	err := yaml.Unmarshal(serviceConfigYamlData, &serviceConfig)
	if err != nil {
		return err
	}

	if serviceConfig.Name == "" || !nameRegex.MatchString(serviceConfig.Name) {
		return fmt.Errorf("service name is invalid")
	}

	if serviceConfig.DisplayName == "" {
		serviceConfig.DisplayName = serviceConfig.Name
	}

	serviceConfig.Describe = cleanstringutils.GetStringOneLine(serviceConfig.Describe)

	switch serviceConfig.ArgumentFrom {
	case FromInstall:
		if len(args) > 0 {
			serviceConfig.ArgumentFrom = FromConfig
			serviceConfig.ArgumentList = sliceutils.CopySlice(args)
		} else {
			serviceConfig.ArgumentFrom = FromNo
			serviceConfig.ArgumentList = nil
		}
	case FromConfig:
		if len(serviceConfig.ArgumentList) == 0 {
			serviceConfig.ArgumentFrom = FromNo
			serviceConfig.ArgumentList = nil
		}
	default:
		serviceConfig.ArgumentFrom = FromNo
		serviceConfig.ArgumentList = nil
	}

	switch serviceConfig.EnvFrom {
	case FromInstall:
		if len(serviceConfig.EnvGetList) == 0 {
			serviceConfig.EnvFrom = FromNo
			serviceConfig.EnvGetList = nil
			serviceConfig.EnvSetList = nil
			break
		}

		serviceConfig.EnvSetList = make(map[string]string, len(serviceConfig.EnvGetList))
		for _, e := range serviceConfig.EnvGetList {
			serviceConfig.EnvSetList[e] = os.Getenv(e)
		}
	case FromConfig:
		serviceConfig.EnvGetList = nil
		if len(serviceConfig.EnvSetList) == 0 {
			serviceConfig.EnvFrom = FromNo
			serviceConfig.EnvSetList = nil
		}
	default:
		serviceConfig.EnvFrom = FromNo
		serviceConfig.EnvGetList = nil
		serviceConfig.EnvSetList = nil
	}

	return nil
}

func initServiceConfig() error {
	err := yaml.Unmarshal(serviceConfigYamlData, &serviceConfig)
	if err != nil {
		return err
	}

	if serviceConfig.Name == "" || !nameRegex.MatchString(serviceConfig.Name) {
		return fmt.Errorf("service name is invalid")
	}

	if serviceConfig.DisplayName == "" {
		serviceConfig.DisplayName = serviceConfig.Name
	}

	serviceConfig.Describe = cleanstringutils.GetStringOneLine(serviceConfig.Describe)

	serviceConfig.ArgumentFrom = FromNo
	serviceConfig.ArgumentList = nil

	serviceConfig.EnvFrom = FromNo
	serviceConfig.EnvGetList = nil
	serviceConfig.EnvSetList = nil

	return nil
}
