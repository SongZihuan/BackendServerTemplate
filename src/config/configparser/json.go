// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"encoding/json"
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

type JsonProvider struct {
	viper   *viper.Viper
	hasRead bool
}

func NewJsonProvider(opt *NewProviderOption) *JsonProvider {
	if opt == nil {
		opt = new(NewProviderOption)
	}

	if opt.EnvPrefix == "" {
		opt.EnvPrefix = envutils.StringToEnvName(osutils.GetArgs0NamePOSIX())
	}

	p := &JsonProvider{
		viper:   viper.New(),
		hasRead: false,
	}

	// 环境变量
	p.viper.SetEnvPrefix(opt.EnvPrefix)
	p.viper.SetEnvKeyReplacer(envutils.GetEnvReplaced())
	p.viper.AutomaticEnv()

	if opt.AutoReload {
		logger.Infof("start auto reload")

		p.viper.OnConfigChange(func(e fsnotify.Event) {
			logger.Infof("config change")

			err := restart.RestartProgram(root.RestartFlag)
			if err != nil {
				logger.Errorf("restart program error: %s", err.Error())
			}
		})

		p.viper.WatchConfig()
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

func (j *JsonProvider) WriteFile(filepath string, src any) configerror.Error {
	if !j.hasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	target, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		return configerror.NewErrorf("json marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, target, 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}

func _testJson() {
	var a ConfigParserProvider
	a = &JsonProvider{}
	_ = a
}
