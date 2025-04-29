// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exitutils

import "testing"

func TestExitCode(t *testing.T) {
	if res := ExitCode(0).ClampAttribute(); res != 0 {
		t.Errorf("ClampAttribute Error: ExitCode(0) -> 0: %d", res)
	}

	if res := ExitCode(1).ClampAttribute(); res != 1 {
		t.Errorf("ClampAttribute Error: ExitCode(1) -> 1: %d", res)
	}

	if res := ExitCode(-5).ClampAttribute(); res != 5 {
		t.Errorf("ClampAttribute Error: ExitCode(-5) -> 5: %d", res)
	}

	if res := ExitCode(300).ClampAttribute(); res != 255 {
		t.Errorf("ClampAttribute Error: ExitCode(300) -> 255: %d", res)
	}

	if res := ExitCode(-400).ClampAttribute(); res != 255 {
		t.Errorf("ClampAttribute Error: ExitCode(-400) -> 255: %d", res)
	}
}

func TestGetExitCode(t *testing.T) {
	if res := getExitCode(0); res != 0 {
		t.Errorf("ClampAttribute Error: getExitCode(0) -> 0: %d", res)
	}

	if res := getExitCode(1); res != 1 {
		t.Errorf("ClampAttribute Error: getExitCode(1) -> 1: %d", res)
	}

	if res := getExitCode(1, 0); res != 0 {
		t.Errorf("ClampAttribute Error: getExitCode(1, 0) -> 0: %d", res)
	}

	if res := getExitCode(1, 0, 2); res != 1 {
		t.Errorf("ClampAttribute Error: getExitCode(1, 0, 2) -> 1: %d", res)
	}

	if res := getExitCode(-1); res != 1 {
		t.Errorf("ClampAttribute Error: getExitCode(-1) -> 1: %d", res)
	}

	if res := getExitCode(400); res != 255 {
		t.Errorf("ClampAttribute Error: getExitCode(400) -> 255: %d", res)
	}

	if res := getExitCode(-300); res != 255 {
		t.Errorf("ClampAttribute Error: getExitCode(-300) -> 255: %d", res)
	}

	if res := getExitCode(0, -1); res != 1 {
		t.Errorf("ClampAttribute Error: getExitCode(0, -1) -> 1: %d", res)
	}

	if res := getExitCode(0, 400); res != 255 {
		t.Errorf("ClampAttribute Error: getExitCode(0, 400) -> 255: %d", res)
	}

	if res := getExitCode(0, -300); res != 255 {
		t.Errorf("ClampAttribute Error: getExitCode(0, -300) -> 255: %d", res)
	}
}
