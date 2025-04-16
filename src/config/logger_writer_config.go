// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/combiningwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/datefilewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/filewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write/wrapwriter"
	"io"
	"os"
	"strings"
)

type LoggerWriterConfig struct {
	WriteToStd          string `json:"write-to-std" yaml:"write-to-std"` // stdout stderr all no
	WriteToFile         string `json:"write-to-file" yaml:"write-to-file"`
	WriteToDirWithDate  string `json:"write-to-dir-with-date" yaml:"write-to-dir-with-date"`
	WriteWithDatePrefix string `json:"write-with-date-prefix" yaml:"write-with-date-prefix"`
}

func (d *LoggerWriterConfig) init(filePath string, provider configparser.ConfigParserProvider) (err configerror.Error) {
	return nil
}

func (d *LoggerWriterConfig) setDefault(c *configInfo) (err configerror.Error) {
	d.WriteToStd = strings.ToLower(d.WriteToStd)

	if d.WriteToStd == "" {
		d.WriteToStd = "stderr"
	}

	if d.WriteToDirWithDate != "" && d.WriteWithDatePrefix == "" {
		d.WriteWithDatePrefix = global.Name
	}

	return nil
}

func (d *LoggerWriterConfig) check(c *configInfo) (err configerror.Error) {
	if d.WriteToStd != "stdout" && d.WriteToStd != "stderr" && d.WriteToStd != "no" {
		return configerror.NewErrorf("bad write-to-std")
	}
	return nil
}

func (d *LoggerWriterConfig) process(c *configInfo, setter func(w io.Writer) (io.Writer, error)) (cfgErr configerror.Error) {
	writerList := make([]write.Writer, 0, 10)

	switch d.WriteToStd {
	case "stdout":
		writerList = append(writerList, wrapwriter.WrapToWriter(os.Stdout))
	case "stderr":
		writerList = append(writerList, wrapwriter.WrapToWriter(os.Stderr))
	}

	if d.WriteToFile != "" {
		fileWriter, err := filewriter.NewFileWriter(d.WriteToFile)
		if err != nil {
			return configerror.NewErrorf("new file writer (on %s) error error: %s", d.WriteToFile, err.Error())
		}

		writerList = append(writerList, fileWriter)
	}

	if d.WriteToDirWithDate != "" {
		dateFileWriter, err := datefilewriter.NewDateFileWriter(d.WriteToDirWithDate, d.WriteWithDatePrefix)
		if err != nil {
			return configerror.NewErrorf("new date file writer (on %s) error error: %s", d.WriteToDirWithDate, err.Error())
		}

		writerList = append(writerList, dateFileWriter)
	}

	if len(writerList) == 0 {
		return nil
	} else if len(writerList) == 1 {
		_, err := setter(writerList[0])
		if err != nil {
			return configerror.NewErrorf("set new writer error: %s", err.Error())
		}
	} else {
		combiningWriter := combiningwriter.NewCombiningWriter(writerList...)
		_, err := setter(combiningWriter)
		if err != nil {
			return configerror.NewErrorf("set new combining writer error: %s", err.Error())
		}
	}

	return nil
}
