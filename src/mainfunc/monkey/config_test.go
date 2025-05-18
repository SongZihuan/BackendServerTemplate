// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monkey

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/global/bddata/bdmodule"
	"strings"
	"testing"
)

// testNormalServiceConfig 此函数为公共函数，并非测试函数（不以Test开头）
func testNormalServiceConfig(t *testing.T, cfg *bdmodule.ServiceConfig) {
	if cfg.DisplayName == "" {
		t.Errorf("Service display name is empty")
	}

	if strings.Contains(cfg.Describe, "\n") {
		t.Errorf("Service describe has LF")
	}

	if strings.Contains(cfg.Describe, "\r") {
		t.Errorf("Service describe has CR")
	}

	if len(cfg.Describe) > 1 && cfg.Describe[0] == ' ' {
		t.Errorf("Service describe start with space")
	}

	if len(cfg.Describe) > 2 && cfg.Describe[len(cfg.Describe)-1] == ' ' {
		t.Errorf("Service describe end with space")
	}
}

func TestNormalServiceConfig(t *testing.T) {
	cfg, err := getRunConfig()
	if err != nil {
		t.Fatalf("Init service error: %s", err.Error())
	}

	testNormalServiceConfig(t, cfg)

	if cfg.ArgumentFrom != bdmodule.FromNo {
		t.Errorf("ArgumentFrom should be no")
	} else if len(cfg.ArgumentList) != 0 {
		t.Errorf("ArgumentList should be empty")
	}

	if cfg.EnvFrom != bdmodule.FromNo {
		t.Errorf("EnvFrom should be no")
	} else if len(cfg.EnvSetList) != 0 {
		t.Errorf("EnvSetList should be empty")
	}
}

func TestInstallServiceConfig(t *testing.T) {
	cfg, err := getRunConfig()
	if err != nil {
		t.Fatalf("Init service error: %s", err.Error())
	}

	testNormalServiceConfig(t, cfg)

	switch cfg.ArgumentFrom {
	case bdmodule.FromConfig:
		if len(cfg.ArgumentList) == 0 {
			t.Errorf("Error argument-from: argument-list is empty, but argument-from config")
		}
	case bdmodule.FromNo:
		if len(cfg.ArgumentList) > 0 {
			t.Errorf("Error argument-from: argument-list is not empty, but argument-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", cfg.ArgumentFrom))
	}

	switch cfg.EnvFrom {
	case bdmodule.FromConfig:
		if len(cfg.EnvSetList) == 0 {
			t.Errorf("Error env-from: env-set-list is empty, but env-from config")
		}
	case bdmodule.FromNo:
		if len(cfg.EnvSetList) > 0 {
			t.Errorf("Error env-from: env-set-list is not empty, but env-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", cfg.EnvFrom))
	}
}
