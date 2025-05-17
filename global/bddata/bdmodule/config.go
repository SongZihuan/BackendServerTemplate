// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bdmodule

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
)

type BuildConfigSet map[string]*BuildConfigData

type BuildConfigData struct {
	Name      *string `yaml:"name"`
	AutoName  *bool   `yaml:"auto-name"`
	EnvPrefix string  `yaml:"env-prefix"`

	Function   FunctionConfig   `yaml:"function"`
	Components ComponentsConfig `yaml:"components"`
	Service    *ServiceConfig   `yaml:"service"`
}

func (b *BuildConfigData) SetDefault(packageName string) error {
	if b.AutoName != nil && b.Name != nil {
		return fmt.Errorf("cannot specify a name and use AutoName at the same time")
	}

	if b.EnvPrefix != "" {
		newEnvPrefix := envutils.ToEnvName(cleanstringutils.GetName(b.EnvPrefix))
		if newEnvPrefix != b.EnvPrefix {
			return fmt.Errorf("env prefix error: use %s please", newEnvPrefix)
		}
	}

	err := b.Function.SetDefault(b)
	if err != nil {
		return err
	}

	err = b.Components.SetDefault(b)
	if err != nil {
		return err
	}

	if b.Service != nil {
		err = b.Service.SetDefault(b)
		if err != nil {
			return err
		}
	}

	return nil
}

type FunctionConfig struct {
	LongVersionInfo  *bool `yaml:"long-version-info"`
	LongVersion      *bool `yaml:"long-version"`
	ShortVersionInfo *bool `yaml:"short-version-info"`
	ShortVersion     *bool `yaml:"short-version"`

	License     *bool `yaml:"license"`
	Report      *bool `yaml:"report"`
	ConfigCheck *bool `yaml:"config-check"`
	Reload      *bool `yaml:"reload"`
}

func (f *FunctionConfig) SetDefault(*BuildConfigData) error {
	setDefault(f.LongVersionInfo, true)
	setDefault(f.LongVersion, true)
	setDefault(f.ShortVersionInfo, true)
	setDefault(f.ShortVersion, true)

	setDefault(f.License, true)
	setDefault(f.Report, true)
	setDefault(f.ConfigCheck, true)
	setDefault(f.Reload, true)

	return nil
}

type ComponentsConfig struct {
	Logger              *bool `yaml:"logger"`
	ConfigWarn2Error    *bool `yaml:"config-warn-to-error"`
	ConfigWatcher       *bool `yaml:"component-watcher"`
	SignalWatcher       *bool `yaml:"signal-watcher"`
	ConsoleEventWatcher *bool `yaml:"console-event-watcher"`
}

func (c *ComponentsConfig) SetDefault(*BuildConfigData) error {
	setDefault(c.Logger, true)
	setDefault(c.ConfigWarn2Error, false)
	setDefault(c.ConfigWatcher, true)
	setDefault(c.SignalWatcher, true)
	setDefault(c.ConsoleEventWatcher, true)

	return nil
}

const (
	FromNo      = "no"
	FromInstall = "install"
	FromConfig  = "config"
)

type ServiceConfig struct {
	DisplayName  string            `yaml:"display-name,omitempty"`
	Describe     string            `yaml:"describe,omitempty"`
	ArgumentFrom string            `yaml:"argument-from,omitempty"`
	ArgumentList []string          `yaml:"argument-list,omitempty"`
	EnvFrom      string            `yaml:"env-from,omitempty"`
	EnvGetList   []string          `yaml:"env-get-list,omitempty"`
	EnvSetList   map[string]string `yaml:"env-set-list,omitempty"`
}

func (s *ServiceConfig) SetDefault(*BuildConfigData) error {
	return nil
}

func setDefault[T any](val *T, def T) *T {
	if val == nil {
		return &def
	}
	return val
}

func set[T any](val T) *T {
	return &val
}
