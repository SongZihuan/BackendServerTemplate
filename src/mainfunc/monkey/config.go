// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monkey

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/bdmodule"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/runner"
	"github.com/SongZihuan/BackendServerTemplate/global/rtdata"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/sliceutils"
)

func getInstallConfig(args []string) (*bdmodule.ServiceConfig, error) {
	cfg, err := runner.GetConfigService()
	if err != nil {
		return nil, err
	} else if cfg == nil {
		return nil, fmt.Errorf("not service config")
	}

	if cfg.Name == "" {
		cfg.Name = rtdata.GetName()
	}

	if cfg.DisplayName == "" {
		cfg.DisplayName = cfg.Name
	}

	if cfg.ArgumentFrom == bdmodule.FromInstall {
		if len(args) > 0 {
			cfg.ArgumentFrom = bdmodule.FromConfig
			cfg.ArgumentList = sliceutils.CopySlice(args)
		} else {
			cfg.ArgumentFrom = bdmodule.FromNo
			cfg.ArgumentList = nil
		}
	} else if len(args) > 0 {
		return nil, fmt.Errorf("no parameters are allowed: %v", args)
	}

	if cfg.EnvFrom == bdmodule.FromInstall {
		cfg.EnvSetList = make(map[string]string, len(cfg.EnvGetList))
		for _, e := range cfg.EnvGetList {
			cfg.EnvSetList[e] = envutils.GetEnv(runner.GetConfigEnvPrefix(), e)
		}

		cfg.EnvFrom = bdmodule.FromConfig
		cfg.EnvGetList = nil
	}

	return cfg, nil
}

func getRunConfig() (*bdmodule.ServiceConfig, error) {
	cfg, err := runner.GetConfigService()
	if err != nil {
		return nil, err
	} else if cfg == nil {
		return nil, fmt.Errorf("not service config")
	}

	if cfg.Name == "" {
		cfg.Name = rtdata.GetName()
	}

	if cfg.DisplayName == "" {
		cfg.DisplayName = cfg.Name
	}

	cfg.ArgumentFrom = bdmodule.FromNo
	cfg.ArgumentList = nil

	cfg.EnvFrom = bdmodule.FromNo
	cfg.EnvGetList = nil
	cfg.EnvSetList = nil

	return cfg, nil
}
