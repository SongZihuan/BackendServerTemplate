// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package timeutils

import (
	"runtime"
	"testing"
)

func TestLocalLocation(t *testing.T) {
	if res := GetLocalTimezone(); res == nil {
		t.Errorf("GetLocalTimezone error: res is nil")
	} else if res.String() == "Local" {
		t.Errorf("GetLocalTimezone error: res.String() is %s", res.String())
	}
}

func TestLoadLocation(t *testing.T) {
	if res, err := LoadTimezone("UTC"); err != nil {
		t.Errorf("LoadTimezone(UTC) error: %s", err.Error())
	} else if res == nil {
		t.Errorf("LoadTimezone(UTC) error: res is nil")
	} else if res.String() != "UTC" {
		t.Errorf("LoadTimezone(UTC) error: res.String() is %s", res.String())
	}

	if res, err := LoadTimezone("Asia/Shanghai"); err != nil {
		t.Errorf("LoadTimezone(Asia/Shanghai) error: %s", err.Error())
	} else if res == nil {
		t.Errorf("LoadTimezone(Asia/Shanghai) error: res is nil")
	} else if res.String() != "Asia/Shanghai" {
		t.Errorf("LoadTimezone(Asia/Shanghai) error: res.String() is %s", res.String())
	}
}

func TestLoadLocationWin32(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skipf("OS is not windows")
	}

	if res, err := LoadTimezone("China Standard Time"); err != nil {
		t.Errorf("LoadTimezone(China Standard Time) error: %s", err.Error())
	} else if res == nil {
		t.Errorf("LoadTimezone(China Standard Time) error: res is nil")
	} else if res.String() != "Asia/Shanghai" {
		t.Errorf("LoadTimezone(China Standard Time) error: res.String() is %s", res.String())
	}
}
