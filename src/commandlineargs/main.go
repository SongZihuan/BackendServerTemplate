package commandlineargs

import (
	"io"
	"os"
)

func Name() string {
	return CommandLineArgsData.NameData
}

func Help() bool {
	return CommandLineArgsData.Help()
}

func PrintUsage() (int, error) {
	return CommandLineArgsData.PrintUsage()
}

func PrintVersion() (int, error) {
	return CommandLineArgsData.PrintVersion()
}

func PrintLicense() (int, error) {
	return CommandLineArgsData.PrintLicense()
}

func PrintReport() (int, error) {
	return CommandLineArgsData.PrintReport()
}

func PrintLF() (int, error) {
	return CommandLineArgsData.PrintLF()
}

func Version() bool {
	return CommandLineArgsData.Version()
}

func License() bool {
	return CommandLineArgsData.License()
}

func Report() bool {
	return CommandLineArgsData.Report()
}

func ConfigFile() string {
	return CommandLineArgsData.ConfigFile()
}

func OutputConfigFile() string {
	return CommandLineArgsData.OutputConfig()
}

func SetOutput(writer io.Writer) {
	if writer == nil {
		writer = os.Stdout
	}

	CommandLineArgsData.SetOutput(writer)
}
