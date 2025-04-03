package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
)

type ExampleData struct {
}

func (d *ExampleData) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *ExampleData) setDefault(c *configInfo) (err configerror.Error) {
	return nil
}

func (d *ExampleData) check(c *configInfo) (err configerror.Error) {
	return nil
}

func (d *ExampleData) process(c *configInfo) (cfgErr configerror.Error) {
	return nil
}
