// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !windows

package timeutils

import (
	"bytes"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"os"
	"os/exec"
	"time"
)

func LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}

func GetLocalTimezone() *time.Location {
	loc, err := getLocalTimezoneFromEnvTZ()
	if err == nil && loc != nil {
		return loc
	}

	loc, err = getLocalTimezoneFromEtcTimezone()
	if err == nil && loc != nil {
		return loc
	}

	loc, err = getLocalTimeZoneFromTimedatectl()
	if err == nil && loc != nil {
		return loc
	}

	return time.UTC
}

func getLocalTimezoneFromEnvTZ() (*time.Location, error) {
	tz := cleanstringutils.GetStringOneLine(os.Getenv("TZ"))
	if tz == "" {
		return nil, fmt.Errorf("TZ not found")
	}
	return time.LoadLocation(tz)
}

func getLocalTimezoneFromEtcTimezone() (*time.Location, error) {
	dat, err := os.ReadFile("/etc/timezone")
	if err != nil {
		return nil, err
	}

	return time.LoadLocation(cleanstringutils.GetStringOneLine(string(dat)))
}

func getLocalTimeZoneFromTimedatectl() (*time.Location, error) {
	// 定义要执行的命令和参数
	cmd := exec.Command("timedatectl", "show", "--property=Timezone", "--value")

	// 捕获命令的标准输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to execute timedatectl: %w", err)
	}

	return time.LoadLocation(cleanstringutils.GetStringOneLine(out.String()))
}
