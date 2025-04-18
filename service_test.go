// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resource

import (
	"fmt"
	"strings"
	"testing"
)

func TestServiceConfig(t *testing.T) {
	if ServiceConfig.Name == "" {
		t.Errorf("Service name is empty")
	}

	if ServiceConfig.DisplayName == "" {
		t.Errorf("Service display name is empty")
	}

	if strings.Contains(ServiceConfig.Describe, "\n") {
		t.Errorf("Service describe has LF")
	}

	if strings.Contains(ServiceConfig.Describe, "\r") {
		t.Errorf("Service describe has CR")
	}

	if len(ServiceConfig.Describe) > 1 && ServiceConfig.Describe[0] == ' ' {
		t.Errorf("Service describe start with space")
	}

	if len(ServiceConfig.Describe) > 2 && ServiceConfig.Describe[len(ServiceConfig.Describe)-1] == ' ' {
		t.Errorf("Service describe end with space")
	}

	switch ServiceConfig.ArgumentFrom {
	case FromConfig:
		if len(ServiceConfig.ArgumentList) == 0 {
			t.Errorf("Error argument-from: argument-list is empty, but argument-from config")
		}
	case FromNo:
		if len(ServiceConfig.ArgumentList) > 0 {
			t.Errorf("Error argument-from: argument-list is not empty, but argument-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", ServiceConfig.ArgumentFrom))
	}

	switch ServiceConfig.EnvFrom {
	case FromConfig:
		if len(ServiceConfig.EnvSetList) == 0 {
			t.Errorf("Error env-from: env-set-list is empty, but env-from config")
		}
	case FromNo:
		if len(ServiceConfig.EnvSetList) > 0 {
			t.Errorf("Error env-from: env-set-list is not empty, but env-from no")
		}
	default:
		t.Errorf(fmt.Sprintf("Error argument-from: %s", ServiceConfig.EnvFrom))
	}
}
