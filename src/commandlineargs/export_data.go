// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commandlineargs

var CommandLineArgsData CommandLineArgsDataType

func (d *CommandLineArgsDataType) Name() string {
	if !d.isReady() {
		panic("flag not ready")
	}

	return d.NameData
}

func (d *CommandLineArgsDataType) Help() bool {
	if !d.isReady() {
		panic("flag not ready")
	}

	return d.HelpData
}

func (d *CommandLineArgsDataType) Version() bool {
	return getData(d, d.VersionData)
}

func (d *CommandLineArgsDataType) OutputVersion() bool {
	return getData(d, d.OutputVersionData)
}

func (d *CommandLineArgsDataType) License() bool {
	return getData(d, d.LicenseData)
}

func (d *CommandLineArgsDataType) Report() bool {
	return getData(d, d.ReportData)
}

func (d *CommandLineArgsDataType) ConfigFile() string {
	return getData(d, d.ConfigFileData)
}

func (d *CommandLineArgsDataType) OutputConfig() string {
	return getData(d, d.ConfigOutputFileData)
}
