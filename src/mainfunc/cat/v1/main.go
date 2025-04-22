// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/SongZihuan/BackendServerTemplate/src/utils/exitutils"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"os"
)

func MainV1(cmd *cobra.Command, args []string, inputConfigFilePath string) (exitCode error) {
	var err error

	logfile, err := os.OpenFile("C:\\Users\\songz\\Code\\GoProject\\BackendServerTemplate\\test_self\\tmpcat.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logfile.Close()
	}()

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	_ = s.Run()
	return prg.ExitCode()
}

func MainV1Install(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initInstallServiceConfig(args)
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 安装服务
	err = s.Install()
	if err != nil {
		return exitutils.InitFailedError("Service Install", err.Error())
	}

	return exitutils.SuccessExitSimple("Service Install Success")
}

func MainV1UnInstall(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 卸载服务
	err = s.Uninstall()
	if err != nil {
		return exitutils.InitFailedError("Service Remove", err.Error())
	}

	return exitutils.SuccessExitSimple("Service Remove Success")
}

func MainV1Start(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 启动服务
	err = s.Start()
	if err != nil {
		return exitutils.InitFailedError("Service Start", err.Error())
	}

	return exitutils.SuccessExitSimple("Service Start Success")
}

func MainV1Stop(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 停止服务
	err = s.Stop()
	if err != nil {
		return exitutils.InitFailedError("Service Stop", err.Error())
	}

	return exitutils.SuccessExitSimple("Service Stop Success")
}

func MainV1Restart(cmd *cobra.Command, args []string) (exitCode error) {
	var err error

	err = initServiceConfig()
	if err != nil {
		return exitutils.InitFailedError("service config", err.Error())
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
		return exitutils.InitFailedError("Service New", err.Error())
	}

	// 重启服务
	err = s.Restart()
	if err != nil {
		return exitutils.InitFailedError("Service Restart", err.Error())
	}

	return exitutils.SuccessExitSimple("Service Restart Success")
}
