// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/restart"
	"github.com/SongZihuan/BackendServerTemplate/utils/envutils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reflect"
	"sync"
)

type JsonProvider struct {
	viper      *viper.Viper
	hasRead    bool
	autoReload bool
	restart    sync.Once
}

func NewJsonProvider(opt *NewConfigParserProviderOption) *JsonProvider {
	if opt == nil {
		opt = new(NewConfigParserProviderOption)
	}

	p := &JsonProvider{
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

func (j *JsonProvider) CanUTF8() bool {
	return true
}

func (j *JsonProvider) ReadFile(filepath string) configerror.Error {
	if j.hasRead {
		return configerror.NewErrorf("config file has been read")
	}

	j.viper.SetConfigFile(filepath)
	j.viper.SetConfigType("json")
	err := j.viper.ReadInConfig()
	if err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return configerror.NewErrorf("config file not found: %s", err.Error())
		}
		return configerror.NewErrorf("read config file error: %s", err.Error())
	}

	if j.autoReload {
		logger.Infof("auto reload: watch file: %s", j.viper.ConfigFileUsed())
		j.viper.WatchConfig()
	}

	j.hasRead = true

	return nil
}

func (j *JsonProvider) ParserFile(target any) configerror.Error {
	if !j.hasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(target).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	err := j.viper.Unmarshal(target)
	if err != nil {
		return configerror.NewErrorf("yaml unmarshal error: %s", err.Error())
	}

	return nil
}
