// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

package consoleutils

import (
	"fmt"
	"os"
	"syscall"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// 获取 FreeConsole 和 AllocConsole 函数
	freeConsole           = kernel32.NewProc("FreeConsole")
	allocConsole          = kernel32.NewProc("AllocConsole")
	setConsoleCtrlHandler = kernel32.NewProc("SetConsoleCtrlHandler")
	getConsoleWindow      = kernel32.NewProc("GetConsoleWindow")
	setConsoleCP          = kernel32.NewProc("SetConsoleCP")
	setConsoleOutputCP    = kernel32.NewProc("SetConsoleOutputCP")
)

func FreeConsole() error {
	ret, _, _ := freeConsole.Call()
	if ret == 0 {
		return fmt.Errorf("FreeConsole error")
	}
	return nil
}

func AllocConsole() error {
	ret, _, _ := allocConsole.Call()
	if ret == 0 {
		return fmt.Errorf("AllocConsole error")
	}
	return nil
}

func BindStdToConsole() error {
	conin, err := os.OpenFile("CONIN$", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	conout, err := os.OpenFile("CONOUT$", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	// 不用关闭旧的标准输入/出/错误
	os.Stdin = conin
	os.Stdout = conout
	os.Stderr = conout

	return nil
}

func SetConsoleCtrlHandler(handler func(event uint) bool, add bool) error {
	var _add uintptr = 0
	if add {
		_add = 1
	}

	ret, _, _ := setConsoleCtrlHandler.Call(
		syscall.NewCallback(func(event uint) uintptr {
			if handler(event) {
				return 1
			}
			return 0
		}),
		_add,
	)
	if ret == 0 {
		return fmt.Errorf("SetConsoleCtrlHandler error")
	}

	return nil
}

func MakeNewConsole(codePage uint) error {
	err := FreeConsole()
	if err != nil {
		return err
	}

	err = AllocConsole()
	if err != nil {
		return err
	}

	err = BindStdToConsole()
	if err != nil {
		return err
	}

	err = SetConsoleCP(codePage)
	if err != nil {
		return err
	}

	return nil
}

func GetConsoleWindow() uint {
	handle, _, _ := getConsoleWindow.Call()
	return uint(handle)
}

func HasConsoleWindow() bool {
	return GetConsoleWindow() != 0
}

func SetConsoleInputCP(codePage uint) error {
	ret, _, _ := setConsoleCP.Call(uintptr(codePage))
	if ret == 0 {
		return fmt.Errorf("SetConsoleInputCP error")
	}
	return nil
}

func SetConsoleOutputCP(codePage uint) error {
	ret, _, _ := setConsoleOutputCP.Call(uintptr(codePage))
	if ret == 0 {
		return fmt.Errorf("SetConsoleOutputCP error")
	}
	return nil
}

func SetConsoleCP(codePage uint) error {
	err := SetConsoleInputCP(codePage)
	if err != nil {
		return err
	}

	err = SetConsoleOutputCP(codePage)
	if err != nil {
		return err
	}

	return nil
}

func SetConsoleCPSafe(codePage uint) error {
	if !HasConsoleWindow() {
		return nil
	}

	return SetConsoleCP(codePage)
}
