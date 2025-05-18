// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monkey

import (
	"github.com/SongZihuan/BackendServerTemplate/src/logger"
	"github.com/SongZihuan/BackendServerTemplate/utils/exitutils"
	"github.com/SongZihuan/BackendServerTemplate/utils/sliceutils"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"strings"
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
		return exitutils.InitFailed("system service init", err.Error())
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

	if svcConfig.Name != svcConfig.DisplayName {
		logger.Infof("Service Name: %s", svcConfig.Name)
		logger.Infof("Service Display Name: %s", svcConfig.DisplayName)
	} else {
		logger.Infof("Service Name & Display Name: %s", svcConfig.Name)
	}

	if svcConfig.Description != "" {
		logger.Infof("Service Description: %s", svcConfig.Description)
	}

	if len(svcConfig.Arguments) == 0 {
		logger.Infof("Service Arguments List: <nil>")
	} else {
		logger.Infof("Service Arguments List: %s", strings.Join(svcConfig.Arguments, " "))
	}

	if len(svcConfig.EnvVars) == 0 {
		logger.Infof("Service EnvVars List: <nil>")
	} else {
		logger.Infof("Service EnvVars List: %s", strings.Join(sliceutils.MapToSlice(svcConfig.EnvVars), " "))
	}

	prg := NewProgram()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return exitutils.InitFailed("system service init", err.Error())
	}

	// 安装服务
	err = s.Install()
	if err != nil {
		return exitutils.InitFailed("system service install", err.Error())
	}

	return exitutils.SuccessExit("system service install success")
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
		return exitutils.InitFailed("system service init", err.Error())
	}

	// 卸载服务
	err = s.Uninstall()
	if err != nil {
		return exitutils.InitFailed("system service remove", err.Error())
	}

	return exitutils.SuccessExit("system service remove success")
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
		return exitutils.InitFailed("system service init", err.Error())
	}

	// 启动服务
	err = s.Start()
	if err != nil {
		return exitutils.InitFailed("system service start", err.Error())
	}

	return exitutils.SuccessExit("system service start success")
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
		return exitutils.InitFailed("system service init", err.Error())
	}

	// 停止服务
	err = s.Stop()
	if err != nil {
		return exitutils.InitFailed("system service stop", err.Error())
	}

	return exitutils.SuccessExit("system service stop success")
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
		return exitutils.InitFailed("system service init", err.Error())
	}

	// 重启服务
	err = s.Restart()
	if err != nil {
		return exitutils.InitFailed("system service restart", err.Error())
	}

	return exitutils.SuccessExit("system service restart success")
}
