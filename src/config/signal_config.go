package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/typeutils"
)

type SignalConfig struct {
	SigIntExit  typeutils.StringBool `json:"sigint-exit" yaml:"sigint-exit"`
	SigTermExit typeutils.StringBool `json:"sigterm-exit" yaml:"sigterm-exit"`
	SigHupExit  typeutils.StringBool `json:"sighup-exit" yaml:"sighup-exit"`
	SigQuitExit typeutils.StringBool `json:"sigquit-exit" yaml:"sigquit-exit"`
}

func (d *SignalConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *SignalConfig) setDefault(c *configInfo) (err configerror.Error) {
	d.SigIntExit.SetDefaultEnable()
	d.SigTermExit.SetDefaultEnable()
	d.SigHupExit.SetDefaultEnable()

	if c.data.IsRelease() {
		d.SigIntExit.SetDefaultDisable()
	} else {
		d.SigIntExit.SetDefaultEnable()
	}

	return nil
}

func (d *SignalConfig) check(c *configInfo) (err configerror.Error) {
	return nil
}

func (d *SignalConfig) process(c *configInfo) (cfgErr configerror.Error) {
	return nil
}
