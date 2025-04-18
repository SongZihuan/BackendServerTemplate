# 变更日志 - Changelog

本项目所有显著变更都将记录在此文件中。

其格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)，
且本项目遵循 [语义化版本控制](https://semver.org/lang/zh-CN/)。

**注意：本文档内容若与[GitHub Wiki](https://github.com/SongZihuan/BackendServerTemplate/wiki/%E5%8F%98%E6%9B%B4%E6%97%A5%E5%BF%97)冲突，则以后者为准**

## [未发布]

### 新增

- 添加对`Github Dependabot` 的支持。

### 修复

- 修复`strconvutils.ReadTimeDuration`中把`uint`转换程`int`可能带来的风险问题，并新增`ReadTimeDurationPositive`函数。
- 完善`README.md`文档关于版本号的描述。

## [0.5.0] - 2025-04-19

### 新增

- 新增`SECURITY.md`、`CONTRIBUTORS.md`和`CONTRIBUTING.md`。
- 为一些文档添加了`Wiki`引用。

### 变更

- 简单修改了一下`README.md`。
- 提供了参考性的`MIT LICENSE`翻译。
- 为部分文件添加遗漏的版权声明。
- 服务描述限制为一行。
- 语义化版本号中的构建信息部分，使用`.`作为分隔符（原使用`-`作为分隔符）。

## [0.4.6] - 2025-04-19

### 新增

- 修复了 0.4.x 系列更新日期问题。
- 补全了 0.4.x 系列更新的日志缺失。

## [0.4.5] - 2025-04-19

### 新增

- 修复了 Github Action 无权限生成 Release 的问题。

## [0.4.4] - 2025-04-19

### 新增

- 添加 GitHub 的 Issue 和 Pull Request 模板。
- Github Action 添加对 Windows的编译支持。

### 变更

- 修改 Github Action 的一项流水线的名字。
- PR合并不再触发 Github Action。

### 修复

- 修复 Windows 在 Github Action 中的环境变量问题。

## [0.4.3] - 2025-04-19

### 修复

- 修复 GitHub Action 生成 Release 时的标题问题（标题中版本号原为`refs/tags/v1.0.0`，现在修改为仅包含`v1.0.0`）。

## [0.4.2] - 2025-04-19

### 修复

- 修复 GitHub Action 获取版本号标签的问题（标签原为`refs/tags/v1.0.0`，现在修改为仅包含`v1.0.0`）。

## [0.4.1] - 2025-04-19

### 修复

- 修复编译问题（删除不受支持的 `gcflag` 参数）。

## [0.4.0] - 2025-04-19

### 新增

- 新增单元测试（以后关于测试的代码变化将记录于 “测试” 小节中）。
- 添加 GitHub Action 配置。

### 测试

- 新增`resource`包的测试。

## [0.3.0] - 2025-04-17

### 新增

- 新增版本号获取功能（仅输出版本号，不输出其他任何内容，不以字母v或V开头）。
- 加入对 `Windows Console` 的支持。
- 添加对机器可读日志的支持（`json`格式）。
- 添加对`Windows`服务的支持。

### 修复

- 修复日志记录器中的按日期分割日志文件记录器丢失日志数据的问题。

### 重构

- 优化了人类可读日志的输出格式。
- 部分原生 `panic` 语句改写为 `logger.Panic` 日志记录。
- 优化命令行参数读取。

## [0.2.0] - 2025-04-16

### 新增

- 获取构建时时间
- 获取构建时`Git`信息（若有）：当前`commit hash`、当前最新`tag`（若有）、以及`tag`（若有）对应的`commit hash`（若有）。
- 清洗通过`go:embed`读取的文件：仅保留第一行（某些文件），删除`BOM`，删除`\r`。
- 新增程序案件退出（`exitutils.SuccessExitQuite()`函数）。

### 变更

- 修改语义化版本号获取：
  - 从`VERSION`文件获取（第一优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
  - 从`git`获取最新的`tag`（第二优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
    - 当该`tag`对应的并非当前`commit`时，`tag`会加上`+dev`标签
    - 当该`tag`以`0.`开头时，`tag`会加上`+dev`标签
  - 采用版本号`0.0.0`
    - 若无`commit hash`，则最终版本号为`0.0.0+dev-1744225466`，其中`1744225466`为编译时间戳。
    - 若有`commit hash`，则最终版本号为`0.0.0+1744225466-be8f4ff51e6ed2e01171b38459406dc5dac306ea`，其中`1744225466`为编译时间戳，`be8f4ff51e6ed2e01171b38459406dc5dac306ea`为`commit hash`。
- `Server.Example1`例子更完善，输出更多信息。
- 命令行参数`--version`输出更多信息：版本号、编译时间（UTC和Local）、编译的Go版本号、系统、架构。
- 应用`exitutils.SuccessExitQuite()`函数到命令行参数的阻断执行退出中。

### 修复

- 修复无法读取`tag`对应`commit`值的漏洞。
- 修复`Output Config File`逻辑判断错误
- 修复设置`SigQuitExit`默认动作的错误
- 修复了退出日志的日志等级
- 修复了命令行参数`--report`不会阻断服务运行的错误。
- 修复命令行`Usage`对短参数的前缀使用错误（原：`--c`，现：`-c`）。

### 重构

- 减少`import resource "github.com/SongZihuan/BackendServerTemplate"`的引用。
- 对日志系统进行跳转：错误日志默认输出到`stderr`。
- 在设定情况下，`config`在执行完`setDefault()`函数后，进行配置文件反向输出（若存在输出路径）。
- 删除`log-name`配置项。

### 文档

- 详细更新了`README`文档。

## [0.1.0] - 2025-04-03

### 新增

- 日志（支持投递到标准输出、文件、日期切割的文件、自定义输出、多输出合并）
- 命令行参数（支持`string`、`bool`、`uint`、`int`）
- 配置文件（支持`json`和`yaml`格式，也可以自定义解析器）
- 退出信号量捕获（在`posix`系统上可以使用信号量捕获退出信号，并做清理操作。在`win32`上，命令行的`ctrl+c`也可被捕获，但当程序作为服务在后台运行时，相关停止、重启操作暂未内捕获）
- 全局变量和资源（打包了`Version`、`License`、`Name`、`Report`等变量）
- 服务模式（可使用控制单元启动多服务，或直接启动单服务）

### 删除

- 删除`test_self`文件夹
