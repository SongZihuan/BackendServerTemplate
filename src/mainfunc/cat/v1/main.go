// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
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
		Name:        global.ServiceConfig.Name,
		DisplayName: global.ServiceConfig.DisplayName,
		Description: global.ServiceConfig.Describe,
		Arguments:   global.ServiceConfig.ArgumentList,
		EnvVars:     global.ServiceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 解析命令行参数
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch strings.ToLower(cmd) {
		case global.Args1Install:
			// 安装服务
			err = s.Install()
			if err != nil {
				return exitutils.InitFailedError("Service Install", err.Error())
			}

			return exitutils.SuccessExitSimple("Service Install Success")
		case global.Args1Uninstall1, global.Args1Uninstall2:
			// 卸载服务
			err = s.Uninstall()
			if err != nil {
				return exitutils.InitFailedError("Service Remove", err.Error())
			}

			return exitutils.SuccessExitSimple("Service Remove Success")
		case global.Args1Start:
			// 启动服务
			err = s.Start()
			if err != nil {
				return exitutils.InitFailedError("Service Start", err.Error())
			}

			return exitutils.SuccessExitSimple("Service Start Success")
		case global.Args1Stop:
			// 停止服务
			err = s.Stop()
			if err != nil {
				return exitutils.InitFailedError("Service Stop", err.Error())
			}

			return exitutils.SuccessExitSimple("Service Stop Success")
		case global.Args1Restart:
			// 重启服务
			err = s.Restart()
			if err != nil {
				return exitutils.InitFailedError("Service Restart", err.Error())
			}

			return exitutils.SuccessExitSimple("Service Restart Success")
		default:
			// 正常运行服务
			// pass
		}
	}

	_ = s.Run()
	return prg.ExitCode()
}
