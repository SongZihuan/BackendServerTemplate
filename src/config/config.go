// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
)

type configInfo struct {
	data *ConfigData

	ready      bool
	inputFile  string
	outputFile string
	provider   configparser.ConfigParserProvider
}

func newConfig(inputFilePath string, outputFilePath string, provider configparser.ConfigParserProvider) (*configInfo, configerror.Error) {
	if inputFilePath == "" {
		logger.Panic("config path is empty")
	}

	configFilePath, err := filesystemutils.CleanFilePathAbs(inputFilePath)
	if err != nil {
		return nil, configerror.NewErrorf("change config file path (%s) to abs error: %s", configFilePath, err.Error())
	}

	if provider == nil {
		provider = configparser.NewYamlProvider(nil)
	}

	if !provider.CanUTF8() {
		return nil, configerror.NewErrorf("config file parser provider new support UTF-8")
	}

	data := new(ConfigData)
	dataInitErr := data.init(configFilePath, provider)
	if dataInitErr != nil && dataInitErr.IsError() {
		return nil, dataInitErr
	}

	return &configInfo{
		data: data,

		ready:      false,
		inputFile:  configFilePath,
		outputFile: outputFilePath,
		provider:   provider,
	}, nil
}

func (c *configInfo) init() (err configerror.Error) {
	if c.ready { // 使用IsReady而不是isReady，确保上锁
		return configerror.NewErrorf("config is ready")
	}

	err = c.provider.ReadFile(c.inputFile)
	if err != nil && err.IsError() {
		return err
	}

	err = c.provider.ParserFile(c.data) // c.Data本身就是指针
	if err != nil && err.IsError() {
		return err
	}

	err = c.data.setDefault(c)
	if err != nil && err.IsError() {
		return err
	}

	if c.outputFile != "" {
		err = c.provider.WriteFile(c.outputFile, c.data)
		if err != nil && err.IsError() {
			return err
		}
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

func (c *configInfo) output(filePath string) configerror.Error {
	if !c.ready {
		return configerror.NewErrorf("config is not ready")
	}

	if filePath == "" {
		return configerror.NewErrorf("config output file path is empty")
	}

	err := c.provider.WriteFile(filePath, c.data)
	if err != nil && err.IsError() {
		return err
	}

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

func (c *configInfo) Output(filePath string) configerror.Error {
	return c.output(filePath)
}
