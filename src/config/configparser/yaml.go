// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/envutils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reflect"
	"sync"
)

type YamlProvider struct {
	viper      *viper.Viper
	autoReload bool
	hasRead    bool
	restart    sync.Once
}

func NewYamlProvider(opt *NewConfigParserProviderOption) *YamlProvider {
	if opt == nil {
		opt = new(NewConfigParserProviderOption)
	}

	p := &YamlProvider{
		viper:      viper.New(),
		autoReload: opt.AutoReload,
		hasRead:    false,
	}

	// 环境变量
	p.viper.SetEnvPrefix(opt.EnvPrefix)
	p.viper.SetEnvKeyReplacer(envutils.EnvReplacer)
	p.viper.AutomaticEnv()

	if p.autoReload {
		logger.Infof("start auto reload")

		p.viper.OnConfigChange(func(e fsnotify.Event) {
			logger.Infof("config change")
			p.restart.Do(func() {
				restart.SetRestart()
			})
		})
	}

	return p
}

func (y *YamlProvider) CanUTF8() bool {
	return true
}

func (y *YamlProvider) ReadFile(filepath string) configerror.Error {
	if y.hasRead {
		return configerror.NewErrorf("config file has been read")
	}

	y.viper.SetConfigFile(filepath)
	y.viper.SetConfigType("yaml")
	err := y.viper.ReadInConfig()
	if err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return configerror.NewErrorf("config file not found: %s", err.Error())
		}
		return configerror.NewErrorf("read config file error: %s", err.Error())
	}

	if y.autoReload {
		logger.Infof("auto reload: watch file: %s", y.viper.ConfigFileUsed())
		y.viper.WatchConfig()
	}

	y.hasRead = true

	return nil
}

func (y *YamlProvider) ParserFile(target any) configerror.Error {
	if !y.hasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	err := y.viper.Unmarshal(target)
	if err != nil {
		return configerror.NewErrorf("yaml unmarshal error: %s", err.Error())
	}

	return nil
}

func _testYaml() {
	var a ConfigParserProvider
	a = &YamlProvider{}
	_ = a
}
