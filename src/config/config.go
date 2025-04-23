// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configoutputer"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
	"os"
	"path"
)

type configInfo struct {
	data *ConfigData

	ready     bool
	hasOutput bool

	inputFile  string
	outputFile string

	parserProvider configparser.ConfigParserProvider
	outputProvider configoutputer.ConfigOutputProvider
}

type ConfigOption struct {
	ConfigFilePath string
	OutputFilePath string
	ParserProvider configparser.ConfigParserProvider
	OutputProvider configoutputer.ConfigOutputProvider
}

func (opt *ConfigOption) setDefault() (err error) {
	if opt.ConfigFilePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			logger.Errorf("can not get work directory: %s", err.Error())
			return err
		}

		opt.ConfigFilePath = path.Join(wd, "config.yaml")
	}

	if opt.ParserProvider == nil {
		opt.ParserProvider, err = configparser.NewConfigParserProvider(opt.ConfigFilePath, nil)
		if err != nil {
			return err
		}
	}

	if opt.OutputFilePath != "" && opt.OutputProvider == nil {
		opt.OutputProvider, err = configoutputer.NewConfigOutputProvider(opt.OutputFilePath, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func newConfig(opt *ConfigOption) (*configInfo, configerror.Error) {
	if opt == nil {
		opt = new(ConfigOption)
	}

	err := opt.setDefault()
	if err != nil {
		return nil, configerror.NewErrorf("new config system error: %s", err)
	}

	configFilePath := opt.ConfigFilePath
	parserProvider := opt.ParserProvider
	outputFilePath := opt.OutputFilePath
	outputProvider := opt.OutputProvider

	data := new(ConfigData)
	dataInitErr := data.init(configFilePath, parserProvider)
	if dataInitErr != nil && dataInitErr.IsError() {
		return nil, dataInitErr
	}

	if configFilePath == "" {
		logger.Panic("config path is empty")
	}

	configFilePath, err = filesystemutils.CleanFilePathAbs(configFilePath)
	if err != nil {
		return nil, configerror.NewErrorf("change config file path (%s) to abs error: %s", configFilePath, err.Error())
	}

	if parserProvider == nil {
		parserProvider = configparser.NewYamlProvider(nil)
	}

	if !parserProvider.CanUTF8() {
		return nil, configerror.NewErrorf("config file parser provider new support UTF-8")
	}

	if outputFilePath != "" {
		outputFilePath, err = filesystemutils.CleanFilePathAbs(outputFilePath)
		if err != nil {
			return nil, configerror.NewErrorf("change config file path (%s) to abs error: %s", configFilePath, err.Error())
		}

		if outputProvider == nil {
			outputProvider = configoutputer.NewYamlProvider(nil)
		}

		if !outputProvider.CanUTF8() {
			return nil, configerror.NewErrorf("config file output provider new support UTF-8")
		}
	}

	return &configInfo{
		data: data,

		ready:     false,
		hasOutput: false,

		inputFile:  configFilePath,
		outputFile: outputFilePath,

		parserProvider: parserProvider,
		outputProvider: outputProvider,
	}, nil
}

func (c *configInfo) init() (err configerror.Error) {
	if c.ready { // 使用IsReady而不是isReady，确保上锁
		return configerror.NewErrorf("config is ready")
	}

	if c.parserProvider == nil {
		return configerror.NewErrorf("config parser provider not set")
	}

	err = c.parserProvider.ReadFile(c.inputFile)
	if err != nil && err.IsError() {
		return err
	}

	err = c.parserProvider.ParserFile(c.data) // c.Data本身就是指针
	if err != nil && err.IsError() {
		return err
	}

	err = c.data.setDefault(c)
	if err != nil && err.IsError() {
		return err
	}

	if c.outputFile != "" && c.outputProvider != nil {
		err = c.outputProvider.WriteFile(c.outputFile, c.data)
		if err != nil && err.IsError() {
			return err
		}
		c.hasOutput = true
	}

	err = c.data.check(c)
	if err != nil && err.IsError() {
		return err
	}

	err = c.data.process(c)
	if err != nil && err.IsError() {
		return err
	}

	c.ready = true
	return nil
}

func (c *configInfo) GetData() (*ConfigData, configerror.Error) {
	if !c.ready {
		return nil, configerror.NewErrorf("config is not ready")
	}

	return c.data, nil
}

func (c *configInfo) Data() *ConfigData {
	if !c.ready {
		logger.Panic("config is not ready")
	}

	return c.data
}

func (c *configInfo) OutputPath() string {
	if c.hasOutput && c.outputProvider != nil && c.outputFile != "" {
		return c.outputFile
	}
	return ""
}
