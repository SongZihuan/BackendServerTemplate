// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commandlineargs

import "github.com/SongZihuan/BackendServerTemplate/src/logger"

var commandLineArgsData commandLineArgsDataType

func (d *commandLineArgsDataType) Name() string {
	if !d.isReady() {
		logger.Panic("flag not ready")
	}

	return d.NameData
}

func (d *commandLineArgsDataType) Help() bool {
	if !d.isReady() {
		logger.Panic("flag not ready")
	}

	return d.HelpData
}

func (d *commandLineArgsDataType) Version() bool {
	return getData(d, d.VersionData)
}

func (d *commandLineArgsDataType) OutputVersion() bool {
	return getData(d, d.OutputVersionData)
}

func (d *commandLineArgsDataType) License() bool {
	return getData(d, d.LicenseData)
}

func (d *commandLineArgsDataType) Report() bool {
	return getData(d, d.ReportData)
}

func (d *commandLineArgsDataType) ConfigFile() string {
	return getData(d, d.ConfigFileData)
}

func (d *commandLineArgsDataType) OutputConfig() string {
	return getData(d, d.ConfigOutputFileData)
}
