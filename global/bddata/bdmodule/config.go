// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bdmodule

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/copyutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
)

type BuildConfigSet map[string]*BuildConfigData

func (b BuildConfigSet) CheckAndSetDefault() error {
	for _, cfg := range b {
		err := cfg.CheckAndSetDefault()
		if err != nil {
			return err
		}
	}

	return nil
}

type BuildConfigData struct {
	Name      *string `yaml:"name,omitempty"`
	AutoName  *bool   `yaml:"auto-name,omitempty"`
	EnvPrefix string  `yaml:"env-prefix,omitempty"`

	Service *ServiceConfig `yaml:"service,omitempty"`
}

func (b *BuildConfigData) CheckAndSetDefault() error {
	if b.AutoName != nil && b.Name != nil {
		return fmt.Errorf("cannot specify a name and use AutoName at the same time")
	}

	if b.EnvPrefix != "" {
		newEnvPrefix := envutils.ToEnvName(b.EnvPrefix)
		if newEnvPrefix != b.EnvPrefix {
			return fmt.Errorf("env prefix error: use %s please", newEnvPrefix)
		}
	}

	if b.Service != nil {
		err := b.Service.CheckAndSetDefault()
		if err != nil {
			return err
		}
	}

	return nil
}

//type FunctionConfig struct {
//	LongVersionInfo  *bool `yaml:"long-version-info"`
//	LongVersion      *bool `yaml:"long-version"`
//	ShortVersionInfo *bool `yaml:"short-version-info"`
//	ShortVersion     *bool `yaml:"short-version"`
//
//	License     *bool `yaml:"license"`
//	Report      *bool `yaml:"report"`
//	ConfigCheck *bool `yaml:"config-check"`
//	Reload      *bool `yaml:"reload"`
//}

//type ComponentsConfig struct {
//	Logger              *bool `yaml:"logger"`
//	ConfigWarn2Error    *bool `yaml:"config-warn-to-error"`
//	ConfigWatcher       *bool `yaml:"component-watcher"`
//	SignalWatcher       *bool `yaml:"signal-watcher"`
//	ConsoleEventWatcher *bool `yaml:"console-event-watcher"`
//}

const (
	FromNo      = "no"
	FromInstall = "install"
	FromConfig  = "config"
)

type ServiceConfig struct {
	Name         string            `yaml:"name,omitempty"`
	DisplayName  string            `yaml:"display-name,omitempty"`
	Describe     string            `yaml:"describe,omitempty"`
	ArgumentFrom string            `yaml:"argument-from,omitempty"`
	ArgumentList []string          `yaml:"argument-list,omitempty"`
	EnvFrom      string            `yaml:"env-from,omitempty"`
	EnvGetList   []string          `yaml:"env-get-list,omitempty"`
	EnvSetList   map[string]string `yaml:"env-set-list,omitempty"`
}

func (s *ServiceConfig) Copy() (*ServiceConfig, error) {
	return copyutils.DeepCopy(s)
}

func (s *ServiceConfig) CheckAndSetDefault() error {
	if s.Name != "" {
		if newName := cleanstringutils.GetName(s.Name); newName != s.Name {
			return fmt.Errorf("service name is invalid: use %s please", newName)
		}
	}

	if s.DisplayName == "" {
		s.DisplayName = s.Name
	}

	s.Describe = cleanstringutils.GetStringOneLine(s.Describe)

	if s.ArgumentFrom == "" || (s.ArgumentFrom == FromConfig && len(s.ArgumentList) == 0) {
		s.ArgumentFrom = FromNo
		s.ArgumentList = nil
	}

	if s.EnvFrom == "" || (s.EnvFrom == FromInstall && len(s.EnvGetList) == 0) || (s.EnvFrom == FromConfig && len(s.EnvSetList) == 0) {
		s.EnvFrom = FromNo
		s.EnvSetList = nil
		s.EnvGetList = nil
	}

	if s.EnvFrom != FromInstall {
		s.EnvGetList = nil
	}

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
