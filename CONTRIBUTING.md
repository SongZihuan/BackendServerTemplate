# Golang 后端服务 模板程序 - 贡献指南

首先，感谢您考虑为本项目做出贡献！我们非常欢迎并珍视来自社区的每一份贡献。无论是代码提交、文档更新还是问题报告，您的参与都是对我们最大的支持。

**注意：本文档内容若与[GitHub Wiki](https://github.com/SongZihuan/BackendServerTemplate/wiki/%E8%B4%A1%E7%8C%AE%E6%8C%87%E5%8D%97)冲突，则以后者为准**

## 如何做出贡献

### 修复Bug

如果您发现了一些Bug，建议您先浏览[Issue](https://github.com/SongZihuan/BackendServerTemplate/issues)，看看是否有类似问题，和解决方案。

更加具体的**Bug报告**要求请参考本项目的 [安全策略](SECURITY.md) 。

### 功能添加

如果您为项目添加了新功能，我们非常欢迎。但出于各种原因，新功能获取不能被采纳，我们对这点深表惋惜。不过您仍然可以将本项目进行`Fork`，然后应用您的更新。
或许一段时间后，或许我们能重新采纳您的新功能。

**注意：本项目是模板Golang项目，不具备实际业务，因此不接受数据库系统、缓存系统、HTTP系统服务、TCP系统服务等相关功能的引入。可以参考我另一个由此项目延伸开发案的HTTP服务系统项目模板。**

### 文档等修缮

如果您认为文档有错误，建议您直接提交[Issue](https://github.com/SongZihuan/BackendServerTemplate/issues)，工作人员查明后进行更正。
此处的文档指：根目录下的`README.md`等文档，以及位于项目`Github Wiki`的文档。

### 代码注释修缮

若您认为代码中存在含糊不清的注释或缺失应有注释，你可以直接提交[Issue](https://github.com/SongZihuan/BackendServerTemplate/issues)，或提交直接提交[PR](https://github.com/SongZihuan/BackendServerTemplate/pulls)。

## 代码风格约束

1. 新建的`.go`文件需包含版权声明头部：
    ```go
    // Copyright 2025 BackendServerTemplate Authors. All rights reserved.
    // Use of this source code is governed by a MIT-style
    // license that can be found in the LICENSE file.
    
    package xxx
    ```

2. 请先阅读`Wiki`，明白各包的作用和相互依赖关系，避免造成以来混乱。
3. 本项目仅是一个模板项目，不具与实际业务。
4. 请编写合适的单元测试，并且提交`PR`后由`GitHub Action`进行测试。
