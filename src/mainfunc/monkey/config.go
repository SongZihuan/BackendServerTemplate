// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monkey

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/bdmodule"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/runner"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/sliceutils"
)

var serviceConfig *bdmodule.ServiceConfig

func initInstallServiceConfig(args []string) error {
	serviceConfig = runner.GetConfig().Service
	if serviceConfig == nil {
		return fmt.Errorf("not service config")
	}

	if serviceConfig.Name == "" {
		serviceConfig.Name = rtdata.GetName()
	} else if newName := cleanstringutils.GetName(serviceConfig.Name); newName != serviceConfig.Name {
		return fmt.Errorf("service name is invalid: use %s please", newName)
	}

	if serviceConfig.DisplayName == "" {
		serviceConfig.DisplayName = serviceConfig.Name
	}

	serviceConfig.Describe = cleanstringutils.GetStringOneLine(serviceConfig.Describe)

	switch serviceConfig.ArgumentFrom {
	case bdmodule.FromInstall:
		if len(args) > 0 {
			serviceConfig.ArgumentFrom = bdmodule.FromConfig
			serviceConfig.ArgumentList = sliceutils.CopySlice(args)
		} else {
			serviceConfig.ArgumentFrom = bdmodule.FromNo
			serviceConfig.ArgumentList = nil
		}
	case bdmodule.FromConfig:
		if len(args) > 0 {
			return fmt.Errorf("no parameters are allowed: %v", args)
		}

		if len(serviceConfig.ArgumentList) == 0 {
			serviceConfig.ArgumentFrom = bdmodule.FromNo
			serviceConfig.ArgumentList = nil
		}
	default:
		if len(args) > 0 {
			return fmt.Errorf("no parameters are allowed: %v", args)
		}

		serviceConfig.ArgumentFrom = bdmodule.FromNo
		serviceConfig.ArgumentList = nil
	}

	switch serviceConfig.EnvFrom {
	case bdmodule.FromInstall:
		if len(serviceConfig.EnvGetList) == 0 {
			serviceConfig.EnvFrom = bdmodule.FromNo
			serviceConfig.EnvGetList = nil
			serviceConfig.EnvSetList = nil
			break
		}

		serviceConfig.EnvSetList = make(map[string]string, len(serviceConfig.EnvGetList))
		for _, e := range serviceConfig.EnvGetList {
			serviceConfig.EnvSetList[e] = envutils.GetSysEnv(e)
		}
	case bdmodule.FromConfig:
		serviceConfig.EnvGetList = nil
		if len(serviceConfig.EnvSetList) == 0 {
			serviceConfig.EnvFrom = bdmodule.FromNo
			serviceConfig.EnvSetList = nil
		}
	default:
		serviceConfig.EnvFrom = bdmodule.FromNo
		serviceConfig.EnvGetList = nil
		serviceConfig.EnvSetList = nil
	}

	return nil
}

func initServiceConfig() error {
	serviceConfig = runner.GetConfig().Service
	if serviceConfig == nil {
		return fmt.Errorf("not service config")
	}

	if serviceConfig.Name == "" {
		serviceConfig.Name = rtdata.GetName()
	} else if newName := cleanstringutils.GetName(serviceConfig.Name); newName != serviceConfig.Name {
		return fmt.Errorf("service name is invalid: use %s please", newName)
	}

	if serviceConfig.DisplayName == "" {
		serviceConfig.DisplayName = serviceConfig.Name
	}

	serviceConfig.Describe = cleanstringutils.GetStringOneLine(serviceConfig.Describe)

	serviceConfig.ArgumentFrom = bdmodule.FromNo
	serviceConfig.ArgumentList = nil

	serviceConfig.EnvFrom = bdmodule.FromNo
	serviceConfig.EnvGetList = nil
	serviceConfig.EnvSetList = nil

	return nil
}
