# 变更日志 - Changelog

本项目所有显著变更都将记录在此文件中。

其格式基于 [CHANGELOG 准则](./CHANGELOG_SPECIFICATION.md) 。

## [未发布]

### 新增

- `cat` 系统也允许处理控制台信号。
- 完善部分单元测试。

### 修复

- 修复了日志无法输出到文件的问题。

### 重构

- 重构了信号量退出包。
- 重构了控制台退出包（仅限`win32`）。
- 重构了服务模型，使用 `Runner` + `Server` + `ServerCore` 模式。用户只需要关心 `ServerCore`即可。

## [0.14.0] - 2025/04/29 Asia/Shanghai

### 新增

- 添加`utils`包的单元测试。

### 修改

- 原于 `cmd` 包下的 `prerun` 包移动到 `src` 下，重命名为 `lifecycle` 包，并使用 `sync.Once` 确保其只执行一次。

### 重构

- 重命名 `lionv1`、 `tigerv1` 和 `catv1` 为 `lion`、 `tiger` 和 `cat` 。

## [0.13.0] - 2025/04/28 Asia/Shanghai

### 新增

- 添加开发者勾子（提交时自动把文件复制粘贴到`third-party`目录）。
- 新增补丁自动生成工具。

### 修改

- 更新依赖库版本。

### 文档

- 新增第三方依赖库记录文件夹（记录`LICENSE`、`CHANGELOG`等）。

### 其他

- 执行`GitHub Action`流水线时，自动执行补丁生成工具，并把补丁放到`Release`中。当用户想更新时，可以使用补丁进行更新。

## [0.12.0] - 2025/04/25 Asia/Shanghai

### 其他

- 新增`CodeQL流水线`。
- 修复流水线生成`release_info.md`时的一些问题。

## [0.11.0] - 2025/04/25 Asia/Shanghai

### 新增

- 在`Windows`平台上可以使用`Windows`时区信息（最终转换为`IANA`时区信息呈现）。
- 配置文件可以交叉输出了，输入`yaml`配置文件，通过`check`子命令可输出为`json`。
- 新增安静模式，启用后标准输入、标准输出、标准错误输出都会输出到`NUL`或`/dev/null`。
- 强制所以`cmd`下的可执行程序包都必须显示导入`prerun`包。
- 强制`prerun`包必须显示导入`global`包。
- 默认导入`time/tzdata`时区数据包，除非添加`systemtzdata`标签，表示使用操作系统自带的数据包。
- 在 `go gerenate` 中添加发布信息（`release_info.md`），但因为该文件需要被忽略，因此修改为：`release_info.md.ignore`。

### 修改

- 去除运行函数的反向配置文件输出能力，仅`check`子命令可以反向输出配置文件。
- 添加`ENV_PREFIX`文件，用于决定咋爱获取与项目有关的环境变量时的前缀。

### 修复

- 修复`format`中遗漏的测试函数。
- 在`Console API`调用前进行`HasConsole`判断，以避免一些潜在的错误。

### 重构

- 完善时区系统，获取本地时间时可以读取到`IANA`时区信息。
- 重构了退出机制。将`main`程序移到`command`程序，将原本的`os.Exit`替换成`return`一个`ExitCode`，最后在`main`函数在使用`os.Exit`退出程序。有效的解决了以前直接在`main`使用`os.Exit`导致`defer`函数无法释放。
- 将 `go generate` 从原本的 `Shell` 脚本（`.sh`和`.ps1`）换成由 `go run` 直接执行的 `.go` 程序。 同时，数据文件以 `.dat` 作为文件后缀（除特殊的 `VERSION`，`NAME`，`REPORT`，`LICENSE`，`ENV_PREFIX`），并且需要忽略的文件以 `.ignore`  作为后缀。

### 文档

- 添加行为准则。
- 完善`README.md`文档。
- 在贡献者指南中把 《Go 编码最佳实际》 列为参考。
- 添加变更日志准则。
- 完善变更日志。
- 删除对`GitHub Wiki`的引用。

### 其他

- 调整`GitHub`的`PR`模板。
- 使用`GitHub Action`流水线创建`Release`时使用`markdown`文件：`release_info.md`。

## [0.10.0] - 2025-04-23 Asia/Shanghai

### 新增

- 添加线程独占任务功能。

### 重构

- 重构日志输出系统（格式化函数、控制台输出ANSI转义序列）。

## [0.9.0] - 2025-04-23 Asia/Shanghai

### 重构

- 重构了自动重启机制。

### 文档

- 完善文档关于命令行参数的讲解。

## [0.8.0] - 2025-04-22 Asia/Shanghai

### 新增

- 添加配置文件检测器。

### 修复

- 修复了配置文件反向输出的问题。

### 重构

- 修改`utils`包下的命名。
- 使用`viper`重构配置文件读取。
- 支持文件重载后，完整重启系统（用于调试功能）。

## [0.7.0] - 2025-04-21 Asia/Shanghai

### 重构

- 优化了命令行匹配系统。

### 文档

- 完善`README.md`文档关于版本号的描述。

## [0.6.0] - 2025-04-19 Asia/Shanghai

### 修复

- 修复`strconvutils.ReadTimeDuration`中把`uint`转换程`int`可能带来的风险问题，并新增`ReadTimeDurationPositive`函数。

### 重构

- 优化了命令行匹配系统，使其能够支持更复杂的命令行（子命令、标志、参数）。

### 文档

- 完善`README.md`文档关于版本号的描述。

### 其他

- 添加对`Github Dependabot` 的支持。

## [0.5.0] - 2025-04-19 Asia/Shanghai

### 修改

- 服务描述限制为一行。
- 语义化版本号中的构建信息部分，使用`.`作为分隔符（原使用`-`作为分隔符）。

### 文档

- 新增`SECURITY.md`、`CONTRIBUTORS.md`和`CONTRIBUTING.md`。
- 为一些文档添加了`Wiki`引用。
- 简单修改了一下`README.md`。
- 提供了参考性的`MIT LICENSE`翻译。
- 为部分文件添加遗漏的版权声明。

## [0.4.6] - 2025-04-19 Asia/Shanghai

### 文档

- 修复了 0.4.x 系列更新日期问题。
- 补全了 0.4.x 系列更新的日志缺失。

## [0.4.5] - 2025-04-19 Asia/Shanghai

### 其他

- 修复了 Github Action 无权限生成 Release 的问题。

## [0.4.4] - 2025-04-19 Asia/Shanghai

### 其他

- 添加 GitHub 的 Issue 和 Pull Request 模板。
- Github Action 添加对 Windows的编译支持。
- 修复 Windows 在 Github Action 中的环境变量问题。
- 修改 Github Action 的一项流水线的名字。
- PR合并不再触发 Github Action。

## [0.4.3] - 2025-04-19 Asia/Shanghai

### 其他

- 修复 GitHub Action 生成 Release 时的标题问题（标题中版本号原为`refs/tags/v1.0.0`，现在修改为仅包含`v1.0.0`）。

## [0.4.2] - 2025-04-19 Asia/Shanghai

### 其他

- 修复 GitHub Action 获取版本号标签的问题（标签原为`refs/tags/v1.0.0`，现在修改为仅包含`v1.0.0`）。

## [0.4.1] - 2025-04-19 Asia/Shanghai

### 修复

- 修复编译问题（删除不受支持的 `gcflag` 参数）。

## [0.4.0] - 2025-04-19 Asia/Shanghai

### 新增

- 新增单元测试（以后关于测试的代码变化将记录于 “测试” 小节中）。
- 添加 GitHub Action 配置。
- 新增`resource`包的测试。

## [0.3.0] - 2025-04-17 Asia/Shanghai

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

## [0.2.0] - 2025-04-16 Asia/Shanghai

### 新增

- 获取构建时时间。
- 获取构建时`Git`信息（若有）：当前`commit hash`、当前最新`tag`（若有）、以及`tag`（若有）对应的`commit hash`（若有）。
- 清洗通过`go:embed`读取的文件：仅保留第一行（某些文件），删除`BOM`，删除`\r`。
- 新增程序案件退出（`exitutils.SuccessExitQuite()`函数）。

### 修改

- 修改语义化版本号获取。
  - 从`VERSION`文件获取（第一优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
  - 从`git`获取最新的`tag`（第二优先级，可以以`v/V`开头，必须满足语义化版本哈规定）。
    - 当该`tag`对应的并非当前`commit`时，`tag`会加上`+dev`标签。
    - 当该`tag`以`0.`开头时，`tag`会加上`+dev`标签。
  - 采用版本号`0.0.0`。
    - 若无`commit hash`，则最终版本号为`0.0.0+dev-1744225466`，其中`1744225466`为编译时间戳。
    - 若有`commit hash`，则最终版本号为`0.0.0+1744225466-be8f4ff51e6ed2e01171b38459406dc5dac306ea`，其中`1744225466`为编译时间戳，`be8f4ff51e6ed2e01171b38459406dc5dac306ea`为`commit hash`。
- `Server.Example1`例子更完善，输出更多信息。
- 命令行参数`--version`输出更多信息：版本号、编译时间（UTC和Local）、编译的Go版本号、系统、架构。
- 应用`exitutils.SuccessExitQuite()`函数到命令行参数的阻断执行退出中。
- 对日志系统进行跳转：错误日志默认输出到`stderr`。
- 删除`log-name`配置项。

### 修复

- 修复无法读取`tag`对应`commit`值的漏洞。
- 修复`Output Config File`逻辑判断错误。
- 修复设置`SigQuitExit`默认动作的错误。
- 修复了退出日志的日志等级。
- 修复了命令行参数`--report`不会阻断服务运行的错误。
- 修复命令行`Usage`对短参数的前缀使用错误（原：`--c`，现：`-c`）。

### 重构

- 减少`import resource "github.com/SongZihuan/BackendServerTemplate"`的引用。
- 在设定情况下，`config`在执行完`setDefault()`函数后，进行配置文件反向输出（若存在输出路径）。

### 文档

- 详细更新了`README.md`文档。

## [0.1.0] - 2025-04-03 Asia/Shanghai

### 新增

- 日志（支持投递到标准输出、文件、日期切割的文件、自定义输出、多输出合并）。
- 命令行参数（支持`string`、`bool`、`uint`、`int`）。
- 配置文件（支持`json`和`yaml`格式，也可以自定义解析器）。
- 退出信号量捕获（在`posix`系统上可以使用信号量捕获退出信号，并做清理操作。在`win32`上，命令行的`ctrl+c`也可被捕获，但当程序作为服务在后台运行时，相关停止、重启操作暂未内捕获）。
- 全局变量和资源（打包了`Version`、`License`、`Name`、`Report`等变量）。
- 服务模式（可使用控制单元启动多服务，或直接启动单服务）。

### 其他

- 删除`test_self`文件夹。
