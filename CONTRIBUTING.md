# Golang 后端服务 模板程序 - 贡献指南

首先，感谢您考虑为本项目做出贡献！我们非常欢迎并珍视来自社区的每一份贡献。无论是代码提交、文档更新还是问题报告，您的参与都是对我们最大的支持。

## 协议与准则

本软件基于 [MIT LICENSE](/LICENSE) 发布，请务必阅读并遵守相关规定。 了解更多关于 MIT LICENSE，请 [点击此处](https://mit-license.song-zh.com) 。

我们欢迎各位贡献者、赞助商等参与项目的开发等各项工作。请务必阅读并遵守我们的 [行为准则](/CODE_OF_CONDUCT.md) 。

同时，欢迎各位贡献者阅读 [贡献者指南](CONTRIBUTING.md) ，有助于我们更好的交流和维护项目。

若您为项目做出了任何贡献，我们将会记录在 [贡献者名单](/CONTRIBUTORS.md) 上，若有任何遗漏可以随时通过 `Issue` 与我们取得联系。

## 如何做出贡献

### 修复Bug

如果您发现了一些Bug，建议您先浏览[Issue](https://github.com/SongZihuan/BackendServerTemplate/issues)，看看是否有类似问题，和解决方案。

更加具体的**Bug报告**要求请参考本项目的 [安全策略](SECURITY.md) 。

### 功能添加

如果您为项目添加了新功能，我们非常欢迎。但出于各种原因，新功能获取不能被采纳，我们对这点深表惋惜。不过您仍然可以将本项目进行`Fork`，然后应用您的更新。
或许一段时间后，或许我们能重新采纳您的新功能。

**注意：本项目是模板Golang项目，不具备实际业务，因此不接受数据库系统、缓存系统、HTTP系统服务、TCP系统服务等相关功能的引入。可以参考我另一个由此项目延伸开发案的HTTP服务系统项目模板。**

### 修改文档

如果您认为文档有错误，建议您直接提交 [Issue](https://github.com/SongZihuan/BackendServerTemplate/issues) ，工作人员查明后进行更正。

### 修改代码注释

若您认为代码中存在含糊不清的注释或缺失应有注释，你可以直接提交 [Issue](https://github.com/SongZihuan/BackendServerTemplate/issues) ，或提交直接提交 [PR](https://github.com/SongZihuan/BackendServerTemplate/pulls) 。

## 代码风格约束

1. 新建的`.go`文件需包含版权声明头部：
    ```go
    // Copyright 2025 BackendServerTemplate Authors. All rights reserved.
    // Use of this source code is governed by a MIT-style
    // license that can be found in the LICENSE file.
    
    package xxx
    ```

2. 本项目仅是一个模板项目，不具与实际业务。
3. 请编写合适的单元测试，并且提交`PR`后由`GitHub Action`进行测试。
4. 代码风格遵守 [《Go 编码最佳实践》](https://google.github.io/styleguide/go/best-practices) 。
