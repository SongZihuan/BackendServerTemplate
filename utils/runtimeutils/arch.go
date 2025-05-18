// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtimeutils

const (
	I386     = "i386"
	AMD64    = "amd64" // x86_64
	ARM      = "arm32"
	ARM64    = "arm64"    // AArch64
	PPC64    = "ppc64"    // PowerPC 64位
	PPC64le  = "ppc64le"  // PowerPC 64位 小端序
	Mips     = "mips"     // MIPS 32位 主要用于嵌入式和旧计算机
	Mipsls   = "mipsls"   // MIPS 32位 小端序 主要用于嵌入式和旧计算机
	Mips64   = "mips64"   // MIPS 64位 主要用于嵌入式和旧计算机
	Mips64ls = "mips64ls" // MIPS 64位 小端序 主要用于嵌入式和旧计算机
	Loong64  = "loong64"  // 龙架构
	S390X    = "s390x"    // 64位 IBM 主机 (System z) 架构，适用于大型机
)

// 推荐用于后端的架构
const (
	ServerArchI386    = I386
	ServerArchAMD64   = AMD64
	ServerArchARM     = ARM
	ServerArchARM64   = ARM64
	ServerArchPPC64   = PPC64
	ServerArchPPC64le = PPC64le
	ServerArchS390X   = S390X
	ServerArchLoong64 = Loong64
)

var ServerArch = map[string]bool{
	ServerArchI386:    true,
	ServerArchAMD64:   true,
	ServerArchARM:     true,
	ServerArchARM64:   true,
	ServerArchPPC64:   true,
	ServerArchPPC64le: true,
	ServerArchS390X:   true,
	ServerArchLoong64: true,
}

var ServerWindowsArch = map[string]bool{
	ServerArchI386:  true,
	ServerArchAMD64: true,
	ServerArchARM:   true,
	ServerArchARM64: true,
}

var ServerLinuxArch = map[string]bool{
	ServerArchI386:    true,
	ServerArchAMD64:   true,
	ServerArchARM:     true,
	ServerArchARM64:   true,
	ServerArchPPC64:   true,
	ServerArchPPC64le: true,
	ServerArchS390X:   true,
	ServerArchLoong64: true,
}

var ServerFreeBSDArch = map[string]bool{
	ServerArchAMD64: true,
	ServerArchARM64: true,
}

var ServerOpenBSDArch = map[string]bool{
	ServerArchAMD64: true,
	ServerArchARM64: true,
}

var ServerNetBSDArch = map[string]bool{
	ServerArchI386:  true,
	ServerArchAMD64: true,
	ServerArchARM:   true,
	ServerArchARM64: true,
}

var ServerOSArch = map[string]map[string]bool{
	ServerOSWindows: ServerWindowsArch,
	ServerOSLinux:   ServerLinuxArch,
	ServerOSFreeBSD: ServerFreeBSDArch,
	ServerOSOpenBSD: ServerOpenBSDArch,
	ServerOSNetBSD:  ServerNetBSDArch,
}
