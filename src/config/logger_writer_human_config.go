// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/datefilewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/filewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/warpwriter"
	"github.com/SongZihuan/BackendServerTemplate/utils/termutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/typeutils"
	"os"
	"strings"
)

type LoggerWriterHumanConfig struct {
	ANSI                typeutils.StringBool `json:"ansi" yaml:"ansi" mapstructure:"ansi"`
	WriteToStd          string               `json:"write-to-std" yaml:"write-to-std" mapstructure:"write-to-std"` // stdout stderr all no
	WriteToFile         string               `json:"write-to-file" yaml:"write-to-file" mapstructure:"write-to-file"`
	WriteToDirWithDate  string               `json:"write-to-dir-with-date" yaml:"write-to-dir-with-date" mapstructure:"write-to-dir-with-date"`
	WriteWithDatePrefix string               `json:"write-with-date-prefix" yaml:"write-with-date-prefix" mapstructure:"write-with-date-prefix"`
}

func (d *LoggerWriterHumanConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *LoggerWriterHumanConfig) setDefault(c *configInfo) (err configerror.Error) {
	d.ANSI.SetDefaultEnable()

	d.WriteToStd = strings.ToLower(d.WriteToStd)

	if d.WriteToDirWithDate != "" && d.WriteWithDatePrefix == "" {
		d.WriteWithDatePrefix = global.Name
	}

	return nil
}

func (d *LoggerWriterHumanConfig) check(c *configInfo) (err configerror.Error) {
	if d.WriteToStd != "stdout" && d.WriteToStd != "stderr" && d.WriteToStd != "no" && d.WriteToStd != "stdout+stderr" && d.WriteToStd != "stderr+stdout" && d.WriteToStd != "" {
		return configerror.NewErrorf("bad write-to-std: %s", d.WriteToStd)
	}
	return nil
}

func (d *LoggerWriterHumanConfig) process(c *configInfo) (writerList []write.Writer, cfgErr configerror.Error) {
	writerList = make([]write.Writer, 0, 10)

	var consoleFn logformat.FormatFunc
	if d.ANSI.IsEnable(true) && termutils.IsTermAdvanced(os.Stdout) && termutils.IsTermAdvanced(os.Stderr) {
		consoleFn = logformat.FormatConsolePretty
	} else {
		consoleFn = logformat.FormatConsole
	}

	switch d.WriteToStd {
	case "stdout":
		writerList = append(writerList, warpwriter.NewWarpWriter(os.Stdout, consoleFn))
	case "stderr":
		writerList = append(writerList, warpwriter.NewWarpWriter(os.Stderr, consoleFn))
	case "stderr+stdout", "stdout+stderr":
		writerList = append(writerList, warpwriter.NewWarpWriter(os.Stdout, consoleFn), warpwriter.NewWarpWriter(os.Stderr, consoleFn))
	case "", "no":
		// pass
	default:
		return nil, configerror.NewErrorf("bad write-to-std: %s", d.WriteToStd)
	}

	if d.WriteToFile != "" {
		fileWriter, err := filewriter.NewFileWriter(d.WriteToFile, logformat.FormatFile)
		if err != nil {
			return nil, configerror.NewErrorf("new file writer (on %s) error error: %s", d.WriteToFile, err.Error())
		}

		writerList = append(writerList, fileWriter)
	}

	if d.WriteToDirWithDate != "" {
		dateFileWriter, err := datefilewriter.NewDateFileWriter(d.WriteToDirWithDate, d.WriteWithDatePrefix, logformat.FormatFile)
		if err != nil {
			return nil, configerror.NewErrorf("new date file writer (on %s) error error: %s", d.WriteToDirWithDate, err.Error())
		}

		writerList = append(writerList, dateFileWriter)
	}

	return writerList, nil
}
