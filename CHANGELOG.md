# 变更日志 - Changelog

本项目所有显著变更都将记录在此文件中。

其格式基于 [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)，
且本项目遵循 [语义化版本控制](https://semver.org/spec/v2.0.0.html)。

## [0.2.0] - 2025-04-10

### 新增功能

- 获取构建时时间
- 获取构建时`Git`信息（若有）：当前`commit hash`、当前最新`tag`（若有）、以及`tag`（若有）对应的`commit hash`（若有）。
- 清洗通过`go:embed`读取的文件：仅保留第一行（某些文件），删除`BOM`，删除`\r`。
- 修改语义化版本号获取：
  - 从`VERSION`文件获取（第一优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
  - 从`git`获取最新的`tag`（第二优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
    - 当该`tag`对应的并非当前`commit`时，`tag`会加上`+dev`标签
    - 当该`tag`以`0.`开头时，`tag`会加上`+dev`标签
  - 采用版本号`0.0.0`
    - 若无`commit hash`，则最终版本号为`0.0.0+dev-1744225466`，其中`1744225466`为编译时间戳。
    - 若有`commit hash`，则最终版本号为`0.0.0+1744225466-be8f4ff51e6ed2e01171b38459406dc5dac306ea`，其中`1744225466`为编译时间戳，`be8f4ff51e6ed2e01171b38459406dc5dac306ea`为`commit hash`。
- `Server.Example1`例子更完善，输出更多信息。

### 修复

- 修复无法读取`tag`对应`commit`值的漏洞。
- 修复`Output Config File`逻辑判断错误
- 修复设置`SigQuitExit`默认动作的错误
- 修复了退出日志的日志等级

### 重构

- 减少`import resource "github.com/SongZihuan/BackendServerTemplate"`的引用。
- 对日志系统进行跳转：错误日志默认输出到`stderr`。

## [0.1.0] - 2025-04-03

### 新增功能

- 日志（支持投递到标准输出、文件、日期切割的文件、自定义输出、多输出合并）
- 命令行参数（支持`string`、`bool`、`uint`、`int`）
- 配置文件（支持`json`和`yaml`格式，也可以自定义解析器）
- 退出信号量捕获（在`posix`系统上可以使用信号量捕获退出信号，并做清理操作。在`win32`上，命令行的`ctrl+c`也可被捕获，但当程序作为服务在后台运行时，相关停止、重启操作暂未内捕获）
- 全局变量和资源（打包了`Version`、`License`、`Name`、`Report`等变量）
- 服务模式（可使用控制单元启动多服务，或直接启动单服务）

### 删除

- 删除`test_self`文件夹
