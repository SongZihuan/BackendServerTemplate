// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

package timeutils

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"os/exec"
	"strings"
	"time"
)

func LoadTimezone(name string) (*time.Location, error) {
	loc, err1 := time.LoadLocation(name)
	if err1 == nil && loc != nil {
		return loc, nil
	}

	loc, err2 := time.LoadLocation(mapWindowsToIANA(name))
	if err2 == nil && loc != nil {
		return loc, nil
	}

	if err1 == nil {
		return nil, fmt.Errorf("location not found")
	}

	return nil, err1
}

func GetLocalTimezone() *time.Location {
	loc, err := getLocalTimezoneFromEnvTZ()
	if err == nil && loc != nil {
		return loc
	}

	loc, err = getLocalTimezoneFromTZutil()
	if err == nil && loc != nil {
		return loc
	}

	return time.UTC
}

func getLocalTimezoneFromEnvTZ() (*time.Location, error) {
	tz := cleanstringutils.GetStringOneLine(envutils.GetSysEnv("TZ"))
	if tz == "" {
		return nil, fmt.Errorf("TZ not found")
	}
	return time.LoadLocation(tz)
}

func getLocalTimezoneFromTZutil() (*time.Location, error) {
	return time.LoadLocation(mapWindowsToIANA(getWindowsTimeZone()))
}

func getWindowsTimeZone() string {
	windowsTZ, err := _getWindowsTimeZone()
	if err != nil {
		return UTC
	}
	return windowsTZ
}

func _getWindowsTimeZone() (string, error) {
	// 使用环境变量或命令行工具获取 Windows 时区名称
	cmd := "tzutil /g"
	out, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}

	windowsTZ := removeDST(cleanstringutils.GetStringOneLine(string(out)))
	_, ok := WindowsTimeZoneMap[windowsTZ]
	if !ok {
		return "", fmt.Errorf("unknown windows timezone: %s", windowsTZ)
	}

	return windowsTZ, nil
}

func removeDST(tz string) string { // 移除夏令时标志
	tz = strings.TrimSuffix(tz, "_dstoff")
	tz = strings.TrimSuffix(tz, "_dston")
	return tz
}

func mapWindowsToIANA(windowsTZ string) string {
	ianaTZ, err := _mapWindowsToIANA(windowsTZ)
	if err != nil {
		return time.UTC.String()
	}

	return ianaTZ
}

func _mapWindowsToIANA(windowsTZ string) (string, error) {
	ianaTZ, exists := WindowsTimeZoneMap[windowsTZ]
	if !exists {
		return "", fmt.Errorf("unknown windows timezone: %s", windowsTZ)
	}
	return ianaTZ, nil
}
