// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtimeutils

// 常见OS
const (
	Windows  = "windows"
	Linux    = "linux"
	Android  = "android"
	IOS      = "ios"
	Darwin   = "darwin"   // MacOS
	FreeBSD  = "freebsd"  // 高性能、文档的Unix
	OpenBSD  = "openbsd"  // 从 NetBSD 分支出来，强调安全性
	NetBSD   = "netbsd"   // 以可移植为主要目标
	DragoFly = "dragofly" // 从 FreeBSD 分离出来，目标是引入新技术和理念
	Solaris  = "solaris"  // SUN 公司开发的 Unix，现在归 Oracle 所有
	Hurd     = "hurd"     // GNU Hurd 操作系统
	JS       = "js"       // 当 Go 程序编译为 JavaScript 在 Web 浏览器中运行时（如使用 GopherJS 或 WebAssembly）
)

// 推荐用于后端的操作系统
const (
	ServerOSWindows = Windows
	ServerOSLinux   = Linux
	ServerOSFreeBSD = FreeBSD
	ServerOSNetBSD  = NetBSD
	ServerOSOpenBSD = OpenBSD
)

var ServerOS = map[string]bool{
	ServerOSWindows: true,
	ServerOSLinux:   true,
	ServerOSFreeBSD: true,
	ServerOSNetBSD:  true,
	ServerOSOpenBSD: true,
}
