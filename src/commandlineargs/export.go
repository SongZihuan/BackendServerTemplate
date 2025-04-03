package commandlineargs

import (
	"fmt"
	"io"
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

func IsReady() bool {
	return CommandLineArgsData.isReady() && isReady
}

func helpInfoRun() error {
	var stopFlag = false

	if Version() {
		_, _ = PrintVersion()
		stopFlag = true
	}

	if License() {
		if stopFlag {
			_, _ = PrintLF()
		}
		_, _ = PrintLicense()
		stopFlag = true
	}

	if Report() {
		if stopFlag {
			_, _ = PrintLF()
		}
		_, _ = PrintReport()
	}

	if Help() {
		if stopFlag {
			_, _ = PrintLF()
		}
		_, _ = PrintUsage()
		stopFlag = true
	}

	if stopFlag {
		return StopRun
	}

	return nil
}
