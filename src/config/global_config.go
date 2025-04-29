// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/timeutils"
	"strings"
	"time"
)

type RunMode string

const (
	RunModeDebug   RunMode = "debug"
	RunModeRelease RunMode = "release"
	RunModeTest    RunMode = "test"
)

type GlobalConfig struct {
	Name     string  `json:"name" yaml:"name" mapstructure:"name"`
	Mode     RunMode `json:"mode" yaml:"mode" mapstructure:"mode"`
	Timezone string  `json:"time-zone" yaml:"time-zone" mapstructure:"time-zone"`

	// Time UTCDate Timestamp 记录为配置文件读取时间
	Time      time.Time `json:"-" yaml:"-" mapstructure:"-"`
	UTCDate   string    `json:"utc-date" yaml:"utc-date" mapstructure:"utc-date"`
	Timestamp int64     `json:"timestamp" yaml:"timestamp" mapstructure:"timestamp"`
}

func (d *GlobalConfig) init(filePath string, provider configparser.ConfigParserProvider) configerror.Error {
	return nil
}

func (d *GlobalConfig) setDefault(c *configInfo) configerror.Error {
	if d.Mode == "" {
		d.Mode = RunModeDebug
	}

	d.Mode = RunMode(strings.ToLower(string(d.Mode)))

	if d.Timezone == "" {
		d.Timezone = "local"
	} else {
		d.Timezone = strings.ToLower(d.Timezone)
	}

	d.Time = time.Now().In(global.UTCLocation)
	d.UTCDate = d.Time.Format(time.DateTime)
	d.Timestamp = d.Time.Unix()

	return nil
}

func (d *GlobalConfig) check(c *configInfo) configerror.Error {
	if d.Mode != RunModeDebug && d.Mode != RunModeRelease && d.Mode != RunModeTest {
		return configerror.NewErrorf("bad mode: %s", d.Mode)
	}

	return nil
}

func (d *GlobalConfig) process(c *configInfo) (cfgErr configerror.Error) {
	name := cleanstringutils.GetStringOneLine(d.Name)
	if (!global.NameFlagChanged || global.Name == "") && name != "" {
		global.Name = name
	}

	var location *time.Location
	if strings.ToLower(d.Timezone) == "utc" {
		location = global.UTCLocation
		if location == nil {
			location = timeutils.GetLocalTimezone()
		}
	} else if strings.ToLower(d.Timezone) == "local" {
		location = timeutils.GetLocalTimezone()
		if location == nil {
			location = global.UTCLocation
		}
	} else {
		var err error
		location, err = timeutils.LoadTimezone(d.Timezone)
		if err != nil || location == nil {
			location = global.UTCLocation
		}

		if location != nil {
			location = timeutils.GetLocalTimezone()
		}
	}

	if location == nil || strings.ToLower(location.String()) == "local" {
		if d.Timezone == "utc" || d.Timezone == "local" {
			return configerror.NewErrorf("can not get location UTC or Local")
		}
		return configerror.NewErrorf("can not get location UTC, Local or %s", d.Timezone)
	}

	global.Location = location

	return nil
}

func (d *GlobalConfig) GetRunMode() RunMode {
	return d.Mode
}

func (d *GlobalConfig) IsDebug() bool {
	return d.Mode == RunModeDebug
}

func (d *GlobalConfig) IsRelease() bool {
	return d.Mode == RunModeRelease
}

func (d *GlobalConfig) IsTest() bool {
	return d.Mode == RunModeTest
}
