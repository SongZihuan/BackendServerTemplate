// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commandlineargs

import (
	"flag"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
)

type CommandLineArgsDataType struct {
	flagReady  bool
	flagSet    bool
	flagParser bool

	NameData  string
	NameName  string
	NameUsage string

	HelpData  bool
	HelpName  string
	HelpUsage string

	VersionData  bool
	VersionName  string
	VersionUsage string

	LicenseData  bool
	LicenseName  string
	LicenseUsage string

	ReportData  bool
	ReportName  string
	ReportUsage string

	ConfigFileData  string
	ConfigFileName  string
	ConfigFileUsage string

	ConfigOutputFileData  string
	ConfigOutputFileName  string
	ConfigOutputFileUsage string

	Usage string
}

func initData() {
	CommandLineArgsData = CommandLineArgsDataType{
		flagReady:  false,
		flagSet:    false,
		flagParser: false,

		NameData:  "", // 默认值为空，具体Name为什么则由config决定
		NameName:  "name",
		NameUsage: fmt.Sprintf("Set the name of the running program, the default is %s.", global.Name),

		HelpData:  false,
		HelpName:  "help",
		HelpUsage: fmt.Sprintf("Show usage of %s. If this option is set, the backend service will not run.", osutils.GetArgs0Name()),

		VersionData:  false,
		VersionName:  "version",
		VersionUsage: fmt.Sprintf("Show version of %s. If this option is set, the backend service will not run.", osutils.GetArgs0Name()),

		LicenseData:  false,
		LicenseName:  "license",
		LicenseUsage: fmt.Sprintf("Show license of %s. If this option is set, the backend service will not run.", osutils.GetArgs0Name()),

		ReportData:  false,
		ReportName:  "report",
		ReportUsage: fmt.Sprintf("Show how to report questions/errors of %s. If this option is set, the backend service will not run.", osutils.GetArgs0Name()),

		ConfigFileData:  "config.yaml",
		ConfigFileName:  "config",
		ConfigFileUsage: fmt.Sprintf("%s", "The location of the running configuration file of the backend service. The option is a string, the default value is config.yaml in the running directory."),

		ConfigOutputFileData:  "",
		ConfigOutputFileName:  "output-config",
		ConfigOutputFileUsage: fmt.Sprintf("%s", "Reverse output configuration file (can be used for corresponding inspection work), the default value is empty (no output)."),

		Usage: "",
	}

	CommandLineArgsData.ready()
}

func (d *CommandLineArgsDataType) setFlag() {
	if d.isFlagSet() {
		return
	}

	flag.StringVar(&d.NameData, d.NameName, d.NameData, d.NameUsage)
	flag.StringVar(&d.NameData, d.NameName[0:1], d.NameData, d.NameUsage)

	flag.BoolVar(&d.HelpData, d.HelpName, d.HelpData, d.HelpUsage)
	flag.BoolVar(&d.HelpData, d.HelpName[0:1], d.HelpData, d.HelpUsage)

	flag.BoolVar(&d.VersionData, d.VersionName, d.VersionData, d.VersionUsage)
	flag.BoolVar(&d.VersionData, d.VersionName[0:1], d.VersionData, d.VersionUsage)

	flag.BoolVar(&d.LicenseData, d.LicenseName, d.LicenseData, d.LicenseUsage)
	flag.BoolVar(&d.LicenseData, d.LicenseName[0:1], d.LicenseData, d.LicenseUsage)

	flag.BoolVar(&d.ReportData, d.ReportName, d.ReportData, d.ReportUsage)
	flag.BoolVar(&d.ReportData, d.ReportName[0:1], d.ReportData, d.ReportUsage)

	flag.StringVar(&d.ConfigFileData, d.ConfigFileName, d.ConfigFileData, d.ConfigFileUsage)
	flag.StringVar(&d.ConfigFileData, d.ConfigFileName[0:1], d.ConfigFileData, d.ConfigFileUsage)

	flag.Usage = func() {
		_, _ = d.PrintUsage()
	}

	d.flagSet = true
}
