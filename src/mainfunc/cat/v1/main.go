// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/consoleutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/kardianos/service"
	"os"
	"strings"
)

func MainV1() (exitCode exitutils.ExitCode) {
	var err error

	err = consoleutils.SetConsoleCPSafe(consoleutils.CodePageUTF8)
	if err != nil {
		return exitutils.InitFailedErrorForWin32ConsoleModule(err.Error())
	}

	err = logger.InitBaseLogger(loglevel.LevelDebug, true, nil, nil, nil, nil)
	if err != nil {
		return exitutils.InitFailedErrorForLoggerModule(err.Error())
	}
	defer logger.CloseLogger()
	defer logger.Recover()

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        fmt.Sprintf("%s-cat-v1", global.Name),
		DisplayName: fmt.Sprintf("%s cat v1", global.Name),
		Description: "简单的Go模板程序",
	}

	// 解析命令行参数
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch strings.ToLower(cmd) {
		case "install":
			if len(os.Args) > 2 {
				svcConfig.Arguments = os.Args[2:]
			}

			prg := NewProgram()
			s, err := service.New(prg, svcConfig)
			if err != nil {
				return exitutils.InitFailedError("Service New", err.Error())
			}

			// 安装服务
			err = s.Install()
			if err != nil {
				return exitutils.InitFailedError("Service Install", err.Error())
			}

			return exitutils.SuccessExit("Service Install Success")
		case "remove", "uninstall":
			prg := NewProgram()
			s, err := service.New(prg, svcConfig)
			if err != nil {
				return exitutils.InitFailedError("Service New", err.Error())
			}

			// 卸载服务
			err = s.Uninstall()
			if err != nil {
				return exitutils.InitFailedError("Service Install", err.Error())
			}

			return exitutils.SuccessExit("Service Install Success")
		case "start":
			prg := NewProgram()
			s, err := service.New(prg, svcConfig)
			if err != nil {
				return exitutils.InitFailedError("Service New", err.Error())
			}

			// 启动服务
			err = s.Start()
			if err != nil {
				return exitutils.InitFailedError("Service Start", err.Error())
			}

			return exitutils.SuccessExit("Service Start Success")
		case "stop":
			prg := NewProgram()
			s, err := service.New(prg, svcConfig)
			if err != nil {
				return exitutils.InitFailedError("Service New", err.Error())
			}

			// 停止服务
			err = s.Stop()
			if err != nil {
				return exitutils.InitFailedError("Service Stop", err.Error())
			}

			return exitutils.SuccessExit("Service Stop Success")
		case "restart":
			prg := NewProgram()
			s, err := service.New(prg, svcConfig)
			if err != nil {
				return exitutils.InitFailedError("Service New", err.Error())
			}

			// 停止服务
			err = s.Restart()
			if err != nil {
				return exitutils.InitFailedError("Service Restart", err.Error())
			}

			return exitutils.SuccessExit("Service Restart Success")
		default:
			// pass
		}
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailedError("Service New", err.Error())
	}

	_ = s.Run()

	return prg.ExitCode()
}
