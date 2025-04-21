// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"strings"
	"testing"
)

func testNormalServiceConfig(t *testing.T) {
	if serviceConfig.Name == "" {
		t.Errorf("Service name is empty")
	}

	if serviceConfig.DisplayName == "" {
		t.Errorf("Service display name is empty")
	}

	if strings.Contains(serviceConfig.Describe, "\n") {
		t.Errorf("Service describe has LF")
	}

	if strings.Contains(serviceConfig.Describe, "\r") {
		t.Errorf("Service describe has CR")
	}

	if len(serviceConfig.Describe) > 1 && serviceConfig.Describe[0] == ' ' {
		t.Errorf("Service describe start with space")
	}

	if len(serviceConfig.Describe) > 2 && serviceConfig.Describe[len(serviceConfig.Describe)-1] == ' ' {
		t.Errorf("Service describe end with space")
	}
}

func TestNormalServiceConfig(t *testing.T) {
	err := initServiceConfig()

	if err != nil {
		t.Fatalf("Init service error: %s", err.Error())
	}

	testNormalServiceConfig(t)

	if serviceConfig.ArgumentFrom != FromNo {
		t.Errorf("ArgumentFrom should be no")
	} else if len(serviceConfig.ArgumentList) != 0 {
		t.Errorf("ArgumentList should be empty")
	}

	if serviceConfig.EnvFrom != FromNo {
		t.Errorf("EnvFrom should be no")
	} else if len(serviceConfig.EnvSetList) != 0 {
		t.Errorf("EnvSetList should be empty")
	}
}

func TestInstallServiceConfig(t *testing.T) {
	err := initServiceConfig()
	if err != nil {
		t.Fatalf("Init service error: %s", err.Error())
	}

	testNormalServiceConfig(t)

	switch serviceConfig.ArgumentFrom {
	case FromConfig:
		if len(serviceConfig.ArgumentList) == 0 {
			t.Errorf("Error argument-from: argument-list is empty, but argument-from config")
		}
	case FromNo:
		if len(serviceConfig.ArgumentList) > 0 {
			t.Errorf("Error argument-from: argument-list is not empty, but argument-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", serviceConfig.ArgumentFrom))
	}

	switch serviceConfig.EnvFrom {
	case FromConfig:
		if len(serviceConfig.EnvSetList) == 0 {
			t.Errorf("Error env-from: env-set-list is empty, but env-from config")
		}
	case FromNo:
		if len(serviceConfig.EnvSetList) > 0 {
			t.Errorf("Error env-from: env-set-list is not empty, but env-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", serviceConfig.EnvFrom))
	}
}
