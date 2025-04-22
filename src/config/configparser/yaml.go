// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configparser

import (
	"errors"
	"github.com/SongZihuan/BackendServerTemplate/src/cmd/restart"
	"github.com/SongZihuan/BackendServerTemplate/src/cmdparser/root"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/envutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"sync"
)

type YamlProvider struct {
	viper      *viper.Viper
	autoReload bool
	reloadLock sync.Mutex
	hasRead    bool
}

func NewYamlProvider(opt *NewProviderOption) *YamlProvider {
	if opt == nil {
		opt = new(NewProviderOption)
	}

	if opt.EnvPrefix == "" {
		opt.EnvPrefix = envutils.StringToEnvName(osutils.GetArgs0NamePOSIX())
	}

	p := &YamlProvider{
		viper:   viper.New(),
		hasRead: false,
	}

	// 环境变量
	p.viper.SetEnvPrefix(opt.EnvPrefix)
	p.viper.SetEnvKeyReplacer(envutils.GetEnvReplaced())
	p.viper.AutomaticEnv()

	if opt.AutoReload {
		logger.Infof("start auto reload")
		p.viper.OnConfigChange(p.reloadEvent)
		p.autoReload = true
	} else {
		p.autoReload = false
	}

	return p
}

func (y *YamlProvider) reloadEvent(e fsnotify.Event) {
	if ok := y.reloadLock.TryLock(); !ok {
		return
	}

	logger.Infof("config change")
	err := restart.RestartProgram(root.RestartFlag)
	if err != nil {
		logger.Errorf("restart program error: %s", err.Error())
		y.reloadLock.Unlock()
		return
	}

	// 不需要释放 y.reloadLock 锁
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

func (y *YamlProvider) WriteFile(filepath string, src any) configerror.Error {
	if !y.hasRead {
		return configerror.NewErrorf("config file has not been read")
	}

	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return configerror.NewErrorf("target must be a pointer")
	}

	target, err := yaml.Marshal(src)
	if err != nil {
		return configerror.NewErrorf("yaml marshal error: %s", err.Error())
	}

	err = os.WriteFile(filepath, target, 0644)
	if err != nil {
		return configerror.NewErrorf("write file error: %s", err.Error())
	}

	return nil
}

func _testYaml() {
	var a ConfigParserProvider
	a = &YamlProvider{}
	_ = a
}
