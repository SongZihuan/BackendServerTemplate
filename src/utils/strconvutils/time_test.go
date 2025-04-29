// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strconvutils

import (
	"testing"
	"time"
)

func TestReadTimeDuration(t *testing.T) {
	if res := ReadTimeDuration("10Min"); res != 10*time.Minute {
		t.Errorf("ReadTimeDuration(10Min) error")
	}

	if res := ReadTimeDuration("-10S"); res != -10*time.Second {
		t.Errorf("ReadTimeDuration(-10S) error")
	}

	if res := ReadTimeDurationPositive("-10S"); res != 0 {
		t.Errorf("ReadTimeDurationPositive(-10S) error")
	}
}

func TestTimeDurationToString(t *testing.T) {
	if res := TimeDurationToString(5 * 24 * time.Hour); res != "5d" {
		t.Errorf("TimeDurationToString(5*24*Hour) -> 5d: %s", res)
	}

	if res := TimeDurationToString(3 * time.Hour); res != "3h" {
		t.Errorf("TimeDurationToString(3*Hour) -> 3h: %s", res)
	}

	if res := TimeDurationToStringCN(5 * 24 * time.Hour); res != "5天" {
		t.Errorf("TimeDurationToString(5*24*Hour) -> 5天: %s", res)
	}

	if res := TimeDurationToStringCN(3 * time.Hour); res != "3小时" {
		t.Errorf("TimeDurationToString(3*Hour) -> 3小时: %s", res)
	}
}
