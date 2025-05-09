// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/SongZihuan/BackendServerTemplate/src/config/configerror"
	"github.com/SongZihuan/BackendServerTemplate/src/config/configparser"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/datefilewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/filewriter"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter/warpwriter"
	"github.com/SongZihuan/BackendServerTemplate/utils/filesystemutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/termutils"
	"io"
	"os"
	"strings"
)

const (
	LoggerWriterTypeStander  = "stander"
	LoggerWriterTypeFile     = "file"
	LoggerWriterTypeDateFile = "date-file"
)

const (
	LoggerFormatTypeConsole          = "console"
	LoggerFormatTypeConsolePretty    = "console-pretty"
	LoggerFormatTypeConsoleTryPretty = "console-try-pretty" // 尝试使用高级颜色字符
	LoggerFormatTypeFile             = "file"
	LoggerFormatTypeJSON             = "json" // 机器模式
)

type LoggerWriterConfig struct {
	Type   string `json:"type"`
	Format string `json:"format" yaml:"format" mapstructure:"format"`

	// 适用于file和date-file类型
	OutputPath string `json:"output-path" yaml:"output-path" mapstructure:"output-path"`

	// 适用于date-file类型
	FilePrefix string `json:"file-prefix" yaml:"file-prefix" mapstructure:"file-prefix"`

	Writer logwriter.Writer `json:"-" yaml:"-" mapstructure:"-"`
}

func (d *LoggerWriterConfig) init(filePath string, provider configparser.ConfigParserProvider) (cfgErr configerror.Error) {
	return nil
}

func (d *LoggerWriterConfig) setDefault(c *configInfo) (cfgErr configerror.Error) {
	if d.Type == "" {
		d.Type = "stander"
	}

	d.Type = strings.ToLower(d.Type)

	if d.Format == "" {
		switch d.Type {
		case LoggerWriterTypeStander:
			d.Format = LoggerFormatTypeConsoleTryPretty
		case LoggerWriterTypeFile, LoggerWriterTypeDateFile:
			d.Format = LoggerFormatTypeFile
		}
	}

	d.Format = strings.ToLower(d.Format)

	if d.OutputPath == "" {
		switch d.Type {
		case LoggerWriterTypeStander:
			d.OutputPath = "stdout"
		case LoggerWriterTypeFile, LoggerWriterTypeDateFile:
			d.OutputPath = "./"
		}
	}

	if d.Type == LoggerWriterTypeStander {
		d.OutputPath = strings.ToLower(d.OutputPath)
	}

	return nil
}

func (d *LoggerWriterConfig) check(c *configInfo) (cfgErr configerror.Error) {
	switch d.Type {
	case LoggerWriterTypeStander:
		if d.OutputPath != "stdout" && d.OutputPath != "stderr" {
			return configerror.NewErrorf("output-path must be 'stdout' or 'stderr': %s", d.OutputPath)
		}
	case LoggerWriterTypeFile:
		if filesystemutils.IsDir(d.OutputPath) {
			return configerror.NewErrorf("output-path is a dir: %s", d.OutputPath)
		}
	case LoggerWriterTypeDateFile:
		if !filesystemutils.IsDir(d.OutputPath) {
			return configerror.NewErrorf("output-path is not a dir: %s", d.OutputPath)
		}
	default:
		return configerror.NewErrorf("error logger type: %s", d.Type)
	}

	if d.Format != LoggerFormatTypeConsole && d.Format != LoggerFormatTypeConsoleTryPretty && d.Format != LoggerFormatTypeConsolePretty &&
		d.Format != LoggerFormatTypeFile && d.Format != LoggerFormatTypeJSON {
		return configerror.NewErrorf("error logger format type: %s", d.Format)

	}

	return nil
}

func (d *LoggerWriterConfig) process(c *configInfo) (writer logwriter.Writer, cfgErr configerror.Error) {
	switch d.Type {
	case LoggerWriterTypeStander:
		var w io.Writer
		if d.OutputPath == "stdout" {
			w = os.Stdout
		} else {
			w = os.Stderr
		}

		switch d.Format {
		case LoggerFormatTypeConsole:
			d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatConsole)
		case LoggerFormatTypeConsolePretty:
			d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatConsolePretty)
		case LoggerFormatTypeConsoleTryPretty:
			if termutils.IsTermAdvanced(w) {
				d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatConsolePretty)
			} else {
				d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatConsole)
			}
		case LoggerFormatTypeFile:
			d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatFile)
		case LoggerFormatTypeJSON:
			d.Writer = warpwriter.NewWarpWriter(w, logformat.FormatJson)
		}
	case LoggerWriterTypeFile:
		switch d.Format {
		case LoggerFormatTypeConsole, LoggerFormatTypeConsoleTryPretty:
			w, err := filewriter.NewFileWriter(d.OutputPath, logformat.FormatConsole)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeConsolePretty:
			w, err := filewriter.NewFileWriter(d.OutputPath, logformat.FormatConsolePretty)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeFile:
			w, err := filewriter.NewFileWriter(d.OutputPath, logformat.FormatFile)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeJSON:
			w, err := filewriter.NewFileWriter(d.OutputPath, logformat.FormatJson)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		}
	case LoggerWriterTypeDateFile:
		switch d.Format {
		case LoggerFormatTypeConsole, LoggerFormatTypeConsoleTryPretty:
			w, err := datefilewriter.NewDateFileWriter(d.OutputPath, d.FilePrefix, logformat.FormatConsole)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeConsolePretty:
			w, err := datefilewriter.NewDateFileWriter(d.OutputPath, d.FilePrefix, logformat.FormatConsolePretty)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeFile:
			w, err := datefilewriter.NewDateFileWriter(d.OutputPath, d.FilePrefix, logformat.FormatFile)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		case LoggerFormatTypeJSON:
			w, err := datefilewriter.NewDateFileWriter(d.OutputPath, d.FilePrefix, logformat.FormatJson)
			if err != nil {
				return nil, configerror.NewErrorf("create filewriter error: %s", err.Error())
			}
			d.Writer = w
		}
	default:
		return nil, configerror.NewErrorf("error logger type: %s", d.Type)
	}

	if d.Writer == nil {
		return nil, configerror.NewErrorf("error logger type: %s", d.Type)
	}

	return d.Writer, nil
}
