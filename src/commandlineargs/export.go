// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commandlineargs

import (
	"fmt"
	"io"
	"os"
)

var StopRun = fmt.Errorf("stop run and exit with success")

var isReady = false

func InitCommandLineArgsParser(output io.Writer) (err error) {
	if IsReady() {
		return nil
	}

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	isReady = false

	initData()

	SetOutput(output)

	isReady = true

	err = helpInfoRun()
	if err != nil {
		return err
	}

	return nil
}

func helpInfoRun() error {
	var stopFlag = false

	if commandLineArgsData.OutputVersion() {
		_, _ = commandLineArgsData.PrintOutputVersion()
		stopFlag = true
		return StopRun
	}

	if commandLineArgsData.Version() {
		_, _ = commandLineArgsData.PrintVersion()
		stopFlag = true
	}

	if commandLineArgsData.License() {
		if stopFlag {
			_, _ = commandLineArgsData.PrintLF()
		}
		_, _ = commandLineArgsData.PrintLicense()
		stopFlag = true
	}

	if commandLineArgsData.Report() {
		if stopFlag {
			_, _ = commandLineArgsData.PrintLF()
		}
		_, _ = commandLineArgsData.PrintReport()
		stopFlag = true
	}

	if commandLineArgsData.Help() {
		if stopFlag {
			_, _ = commandLineArgsData.PrintLF()
		}
		_, _ = commandLineArgsData.PrintUsage()
		stopFlag = true
	}

	if stopFlag {
		return StopRun
	}

	return nil
}

func IsReady() bool {
	return commandLineArgsData.isReady() && isReady
}

func Name() string {
	return commandLineArgsData.Name()
}

func ConfigFile() string {
	return commandLineArgsData.ConfigFile()
}

func OutputConfigFile() string {
	return commandLineArgsData.OutputConfig()
}

func SetOutput(writer io.Writer) {
	if writer == nil {
		writer = os.Stdout
	}

	commandLineArgsData.SetOutput(writer)
}
