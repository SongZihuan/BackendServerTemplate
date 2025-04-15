package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ConfigData struct {
	GlobalConfig `json:",inline" yaml:",inline"`
	Logger       LoggerConfig `json:"logger" yaml:"logger"`
	Signal       SignalConfig `json:"signal" yaml:"signal"`
	Server       ServerConfig `json:"server" yaml:"server"`
}

func (d *ConfigData) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	cfgErr := d.GlobalConfig.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.init(filePath, provider)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) setDefault(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.setDefault(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) check(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.check(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}

func (d *ConfigData) process(c *configInfo) (err configerror.Error) {
	cfgErr := d.GlobalConfig.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Logger.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Signal.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	cfgErr = d.Server.process(c)
	if cfgErr != nil {
		return cfgErr
	}

	return nil
}
