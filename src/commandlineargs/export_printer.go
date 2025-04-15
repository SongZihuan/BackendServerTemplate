// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commandlineargs

import (
	"flag"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/formatutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"io"
	"runtime"
	"strings"
	"time"
)

func (d *CommandLineArgsDataType) FprintUsage(writer io.Writer) (int, error) {
	return fmt.Fprintf(writer, "%s\n", d.Usage)
}

func (d *CommandLineArgsDataType) PrintUsage() (int, error) {
	return d.FprintUsage(flag.CommandLine.Output())
}

func (d *CommandLineArgsDataType) FprintVersion(writer io.Writer) (int, error) {
	res := new(strings.Builder)
	res.WriteString(fmt.Sprintf("Version of %s: %s\n", osutils.GetArgs0Name(), global.Version))
	res.WriteString(fmt.Sprintf("Build Date (UTC): %s\n", global.BuildTime.In(time.UTC).Format(time.DateTime)))
	if time.Local != nil && time.Local.String() != time.UTC.String() {
		zone, _ := time.Now().Local().Zone()
		res.WriteString(fmt.Sprintf("Build Date (%s): %s\n", zone, global.BuildTime.In(time.Local).Format(time.DateTime)))
	}
	res.WriteString(fmt.Sprintf("Compiler: %s\n", runtime.Version()))
	res.WriteString(fmt.Sprintf("OS: %s\n", runtime.GOOS))
	res.WriteString(fmt.Sprintf("Arch: %s\n", runtime.GOARCH))

	version := formatutils.FormatTextToWidth(res.String(), formatutils.NormalConsoleWidth)
	return fmt.Fprintf(writer, "%s\n", version)
}

func (d *CommandLineArgsDataType) PrintVersion() (int, error) {
	return d.FprintVersion(flag.CommandLine.Output())
}

func (d *CommandLineArgsDataType) FprintLicense(writer io.Writer) (int, error) {
	title := formatutils.FormatTextToWidth(fmt.Sprintf("License of %s:", osutils.GetArgs0Name()), formatutils.NormalConsoleWidth)
	license := formatutils.FormatTextToWidth(global.License, formatutils.NormalConsoleWidth)
	return fmt.Fprintf(writer, "%s\n%s\n", title, license)
}

func (d *CommandLineArgsDataType) PrintLicense() (int, error) {
	return d.FprintLicense(flag.CommandLine.Output())
}

func (d *CommandLineArgsDataType) FprintReport(writer io.Writer) (int, error) {
	// 不需要title
	report := formatutils.FormatTextToWidth(global.Report, formatutils.NormalConsoleWidth)
	return fmt.Fprintf(writer, "%s\n", report)
}

func (d *CommandLineArgsDataType) PrintReport() (int, error) {
	return d.FprintReport(flag.CommandLine.Output())
}

func (d *CommandLineArgsDataType) FprintLF(writer io.Writer) (int, error) {
	return fmt.Fprintf(writer, "\n")
}

func (d *CommandLineArgsDataType) PrintLF() (int, error) {
	return d.FprintLF(flag.CommandLine.Output())
}
