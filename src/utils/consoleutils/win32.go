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

const ATTACH_PARENT_PROCESS = uintptr(^uint32(0)) // 0xFFFFFFFF

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// 获取 FreeConsole 和 AllocConsole 函数
	freeConsole           = kernel32.NewProc("FreeConsole")
	allocConsole          = kernel32.NewProc("AllocConsole")
	setConsoleCtrlHandler = kernel32.NewProc(
		"SetConsoleCtrlHandler")
	getConsoleWindow   = kernel32.NewProc("GetConsoleWindow")
	setConsoleCP       = kernel32.NewProc("SetConsoleCP")
	setConsoleOutputCP = kernel32.NewProc("SetConsoleOutputCP")
	attachConsole      = kernel32.NewProc("AttachConsole")
)

func FreeConsole() error {
	ret, _, err := freeConsole.Call()
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("FreeConsole error: %s", err.Error())
	}
	return nil
}

func AllocConsole() error {
	ret, _, err := allocConsole.Call()
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknow")
		}
		return fmt.Errorf("AllocConsole error: %s", err.Error())
	}
	return nil
}

func BindStdToConsoleSafe() error {
	if !HasConsoleWindow() {
		return nil
	}

	return BindStdToConsole()
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

	ret, _, err := setConsoleCtrlHandler.Call(
		syscall.NewCallback(func(event uint) uintptr {
			if handler(event) {
				return 1
			}
			return 0
		}),
		_add,
	)
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("SetConsoleCtrlHandler error: %s", err)
	}

	return nil
}

func MakeNewConsole(codePage uint) error {
	if HasConsoleWindow() {
		err := FreeConsole()
		if err != nil {
			return err
		}
	}

	err := AllocConsole()
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
	ret, _, err := setConsoleCP.Call(uintptr(codePage))
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("SetConsoleInputCP error: %s", err.Error())
	}
	return nil
}

func SetConsoleOutputCP(codePage uint) error {
	ret, _, err := setConsoleOutputCP.Call(uintptr(codePage))
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("SetConsoleOutputCP error: %s", err.Error())
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

func AttachConsole(ppid int) error {
	if HasConsoleWindow() {
		err := FreeConsole()
		if err != nil {
			return err
		}
	}

	// 定义目标进程 ID
	ret, _, err := attachConsole.Call(uintptr(ppid))
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("AttachParentConsole error: %s", err.Error())
	}

	return nil
}

func AttachParentConsole() error {
	if HasConsoleWindow() {
		err := FreeConsole()
		if err != nil {
			return err
		}
	}

	// 定义目标进程 ID
	ret, _, err := attachConsole.Call(ATTACH_PARENT_PROCESS)
	if ret == 0 {
		if err == nil {
			err = fmt.Errorf("unknown")
		}
		return fmt.Errorf("AttachParentConsole error: %s", err.Error())
	}

	return nil
}
