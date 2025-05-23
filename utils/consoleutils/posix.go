// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !windows

package consoleutils

func FreeConsole() error {
	return nil
}

func AllocConsole() error {
	return nil
}

func BindStdToConsole() error {
	return nil
}

func SetConsoleCtrlHandler(handler func(event uint) bool, add bool) error {
	return nil
}

func MakeNewConsole() error {
	return nil
}

func GetConsoleWindow() uintptr {
	return 0
}

func HasConsoleWindow() bool {
	return GetConsoleWindow() != 0
}

func SetConsoleInputCP(codePage uint) error {
	return nil
}

func SetConsoleOutputCP(codePage uint) error {
	return nil
}

func SetConsoleCP(codePage uint) error {
	return nil
}

func SetConsoleCPSafe(codePage uint) error {
	return nil
}

func AttachConsole() error {
	return nil
}
