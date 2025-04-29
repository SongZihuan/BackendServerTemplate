// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cat

import (
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string, inputConfigFilePath string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewRunProgram(inputConfigFilePath)
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	_ = s.Run()
	return prg.ExitCode()
}

func MainInstall(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initInstallServiceConfig(args)
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	// 安装服务
	err = s.Install()
	if err != nil {
		return exitutils.InitFailed("Service Install", err.Error())
	}

	return exitutils.SuccessExit("Service Install Success")
}

func MainUnInstall(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	// 卸载服务
	err = s.Uninstall()
	if err != nil {
		return exitutils.InitFailed("Service Remove", err.Error())
	}

	return exitutils.SuccessExit("Service Remove Success")
}

func MainStart(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	// 启动服务
	err = s.Start()
	if err != nil {
		return exitutils.InitFailed("Service Start", err.Error())
	}

	return exitutils.SuccessExit("Service Start Success")
}

func MainStop(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	// 停止服务
	err = s.Stop()
	if err != nil {
		return exitutils.InitFailed("Service Stop", err.Error())
	}

	return exitutils.SuccessExit("Service Stop Success")
}

func MainRestart(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailed("service config", err.Error())
	}

	// 定义服务配置
	svcConfig := &service.Config{
		Name:        serviceConfig.Name,
		DisplayName: serviceConfig.DisplayName,
		Description: serviceConfig.Describe,
		Arguments:   serviceConfig.ArgumentList,
		EnvVars:     serviceConfig.EnvSetList,
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("Service New", err.Error())
	}

	// 重启服务
	err = s.Restart()
	if err != nil {
		return exitutils.InitFailed("Service Restart", err.Error())
	}

	return exitutils.SuccessExit("Service Restart Success")
}
