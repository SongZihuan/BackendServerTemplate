package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/filesystemutils"
)

type configInfo struct {
	data *ConfigData

	ready    bool
	file     string
	provider configparser.ConfigParserProvider
}

func newConfig(filePath string, provider configparser.ConfigParserProvider) (*configInfo, configerror.Error) {
	if filePath == "" {
		panic("config path is empty")
	}

	configFilePath, err := filesystemutils.CleanFilePathAbs(filePath)
	if err != nil {
		return nil, configerror.NewErrorf("change config file path (%s) to abs error: %s", configFilePath, err.Error())
	}

	if provider == nil {
		provider = configparser.NewYamlProvider()
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

		ready:    false,
		file:     configFilePath,
		provider: provider,
	}, nil
}

func (c *configInfo) init() (err configerror.Error) {
	if c.ready { // 使用IsReady而不是isReady，确保上锁
		return configerror.NewErrorf("config is ready")
	}

	err = c.provider.ReadFile(c.file)
	if err != nil && err.IsError() {
		return err
	}

	err = c.provider.ParserFile(c.data)
	if err != nil && err.IsError() {
		return err
	}

	err = c.data.setDefault(c)
	if err != nil && err.IsError() {
		return err
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
		panic("config is not ready")
	}

	return c.data
}

func (c *configInfo) ConfigFilePath() string {
	return c.file
}

func (c *configInfo) Output(filePath string) configerror.Error {
	return c.output(filePath)
}
