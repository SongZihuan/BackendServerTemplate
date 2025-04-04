package config

import (
	resource "github.com/SongZihuan/BackendServerTemplate"
	"github.com/SongZihuan/BackendServerTemplate/src/commandlineargs"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
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
	Name     string  `json:"name" yaml:"name"`
	Mode     RunMode `json:"mode" yaml:"mode"`
	Timezone string  `json:"time-zone" yaml:"time-zone"`
}

func (d *GlobalConfig) init(filePath string, provider configparser.ConfigParserProvider) configerror.Error {
	return nil
}

func (d *GlobalConfig) setDefault() configerror.Error {
	if d.Mode == "" {
		d.Mode = RunModeDebug
	}

	d.Mode = RunMode(strings.ToLower(string(d.Mode)))

	if d.Timezone == "" {
		d.Timezone = "local"
	} else {
		d.Timezone = strings.ToLower(d.Timezone)
	}

	return nil
}

func (d *GlobalConfig) check() configerror.Error {
	if d.Mode != RunModeDebug && d.Mode != RunModeRelease && d.Mode != RunModeTest {
		return configerror.NewErrorf("bad mode: %s", d.Mode)
	}

	return nil
}

func (d *GlobalConfig) process(c *configInfo) (cfgErr configerror.Error) {
	if commandlineargs.Name() != "" {
		global.Name = commandlineargs.Name()
	} else if d.Name != "" {
		global.Name = d.Name
	} else {
		global.Name = resource.Name
	}

	d.Name = global.Name

	var location *time.Location
	if strings.ToLower(d.Timezone) == "utc" {
		location = time.UTC
		if location == nil {
			location = time.Local
		}
	} else if strings.ToLower(d.Timezone) == "local" {
		location = time.Local
		if location == nil {
			location = time.UTC
		}
	} else {
		var err error
		location, err = time.LoadLocation(d.Timezone)
		if err != nil || location == nil {
			location = time.UTC
		}

		if location != nil {
			location = time.Local
		}
	}

	if location == nil {
		if d.Timezone == "utc" || d.Timezone == "local" {
			return configerror.NewErrorf("can not get location UTC or Local")
		}

		return configerror.NewErrorf("can not get location UTC, Local or %s", d.Timezone)
	} else {
		global.Location = location
	}

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
