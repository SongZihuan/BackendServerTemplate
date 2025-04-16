# Golang 后端服务 模板程序 - 简单版

使用 `Golang` 实现的后端服务模板程序。

## 介绍

本模板程序旨在实现一个 `Golang` 的后端服务，可以是 `Http` 也可以是其他。为了完成这个目的，我设计了一部分基础套件。

* 日志（支持投递到标准输出、文件、日期切割的文件、自定义输出、多输出合并）
* 命令行参数（支持`string`、`bool`、`uint`、`int`）
* 配置文件（支持`json`和`yaml`格式，也可以自定义解析器）
* 退出信号量捕获（在`posix`系统上可以使用信号量捕获退出信号，并做清理操作。在`win32`上，命令行的`ctrl+c`也可被捕获，但当程序作为服务在后台运行时，相关停止、重启操作暂未内捕获）
* 全局变量和资源（打包了`Version`、`License`、`Name`、`Report`等变量）
* 服务模式（可使用控制单元启动多服务，或直接启动单服务）

## 入口

入口文件在`src/cmd`下，目前分别有两个程序：`lionv1`和`tigerv1`。

* `lionv1` 是使用控制单元的多服务演示程序。
* `tigerv1` 是直接运行服务的单服务演示程序。

入口程序不直接包含太多的实际代码，真正的`main`函数位于`src\mainfunc`下。
程序的返回值代表程序的`Exit Code`。
一般采用`0`表示正确结束。

## 编译

### 生成

每次编译器请先运行`go generate`命令，生成编译所需要的文件。

编译所需要的文件：

* `build_data.txt` 构建日期（Unix 时间戳，单位：秒）
* `commit_data.txt` 构建的`git`（若有）的`commit`所对应的`hash`值（完整）。
* `tag_data.txt` 构建时最新的`git`（若有）的标签名（若有），用于作为语义化版本。可以以`v`或`V`开头，后接语义化版本号。若非语义化版本号，则标签数据会忽略（相当于无标签）。
* `tag_commit_data.txt` 上述（`tag_data.txt`）对标签（若有）所指的`commit`的`hash`值（完整）。若该值与`commit_data.txt`不同，则版本号会加上`+dev`标识。
* `random_data.txt` 一串随机数（40个字符，和`commit`的`hash`长度相同，由数字和小写字母组成）。

在项目根目录下执行：
```shell
$ go generate ./...
```

其中，`./...`表示执行当前目录下的所有（含本目录和递归的子目录）`go:generate`命令。

### 编译

使用`go build`指令进行编译。如何包位于`github.com/SongZihuan/BackendServerTemplate/src/cmd/`下，例如`github.com/SongZihuan/BackendServerTemplate/src/cmd/lionv1`

所以，编译命令如下（以`lionv1`为例）：
```shell
$ go build github.com/SongZihuan/BackendServerTemplate/src/cmd/lionv1
```

若用于开发环境，可以按如下方式编译：
```shell
$ go build -o lionv1 -trimpath -ldflags='-s -w -extldflags "-static"' -gcflags='-O4 -inline=1' github.com/SongZihuan/BackendServerTemplate/src/cmd/lionv1
```

其中：
 * `-o lionv1` 表示输出二进制文件的路径和名称
   * Windows： `-o lionv1.exe`
   * Linux/MacOS：`-o lionv1`
 * `-trimpath` 表示擦除二进制文件中关于源码目录的信息，可以压缩二进制文件体积和包含隐私。
 * `-ldflags='-s -w'` 是传递给链接器的参数。
   * `-s` 表示去掉符号和调试信息，可以减少二进制文件大小，同时增加反编译的难度。
   * `-w` 表示去掉`DWARF`调试信息，可以减少二进制文件大小，同时增加反编译的难度。
   * `-extldflags "-static"` 是传递给外部链接器的参数。
     * `-static` 表示外部链接器生成静态链接文件（实际上我几乎从真正看过这个参数的作用，因为`go`本身就会优先以静态形式链接库文件，这个参数可能和`cgo`搭配更合适）。
 * `-gcflags='-O4 -inline=1'` 是传递给编译器的参数。
   * `-O4` 启用最大优化。
   * `-inline=1` 内联的积极程度。支持-1、0、1、2、3。数值越大约积极，-1为禁止内联，0为默认模式。内联越积极，可以提高一定性能，但可能让二进制文件膨胀变大。

## 运行
执行编译好地可执行文件即可。
**注意编译时选择的目标平台要与运行平台一致。**

### 运行参数

* `-h`、`--help` 输出帮助文档（`Usage`）。**（终止运行参数）。**
* `-v`、`--version` 输出版本号。**（终止运行参数）。**
* `-l`、`--license` 输出软件协议。**（终止运行参数）。**
* `-r`、`--report` 输出程序反馈、报告方式文档。**（终止运行参数）。**
* `-n`、`--name` 设置项目名称（需要一个字符串参数）
* `-c`、`--config` 配置文件位置（可提供一个字符串参数作为路径指向配置文件，默认值为：`config.yaml`）。
* `-o`、`--output-config` 反向输出配置文件（默认不输出，若提供一个路径则会输出到该路径所指处）。未知配置项不会输出，未设定配置项则以默认值（若有）输出。**前提：配置文件被正确导入和识别。**

**终止运行参数：当命令行出现这些参数时，将只执行参数对应的功能，执行完成后不会继续运行后续服务。**

### 配置文件

```yaml
# 等同于命令行参数的 --name ，但优先级高于命令行参数。
name: Backend-Server-Template

# 运行模式：debug、release、test。默认为debug。
mode: debug

# 时间地区，例如：UTC、Local（默认）、Asia/Shanghai
time-zone: local

# utc-date 和 timestamp 并非真实参数，而是启用 --output-config时反向输出的参数、表示配置文件的读取时间。
utc-date: "2025-04-15 15:33:37"
timestamp: 1744731217

# 日志记录器的配置
logger:
    log-level: debug  # 日志记录等级：debug（输出debug和以上的） < info （输出info和以上的）< warn < error < panic < none（什么都不输出）
    log-tag: enable  # 是否输出tag调试日志。
    
    human-warn-writer:  # 人类可读的 debug、tag、info、warn 日志的输出器
        write-to-std: stdout  # 输出到标准输出或标准错误输出（为空则不启用）
        write-to-file: ""  # 输出到固定文件（append）模式
        write-to-dir-with-date: ""  # 输出到指定目录，并按日期分割，此处为输出路径
        write-with-date-prefix: ""  # 配合 write-to-dir-with-date ，表示文件的输出前缀
        
    human-error-writer:  # 人类可读的 error、panic 日志的输出器，含义同上
        write-to-std: stdout
        write-to-file: ""
        write-to-dir-with-date: ""
        write-with-date-prefix: ""

    machine-warn-writer:  # 机器可读的 debug、tag、info、warn 日志的输出器，含义同上
      write-to-std: stdout
      write-to-file: ""
      write-to-dir-with-date: ""
      write-with-date-prefix: ""
      
    machine-error-writer:  # 机器可读的 error、panic 日志的输出器，含义同上
      write-to-std: stdout
      write-to-file: ""
      write-to-dir-with-date: ""
      write-with-date-prefix: ""

signal: # 信号除了机制（管理接收程序退出信号）。sigkill 等信号是不可捕获的，是强制退出的，因此此处无法控制这类信号。虽然windows本身不具有Linux这种信号机制，但是Go在信号方面做了一层模拟，使得控制它ctrl+c可以转换为相应信号。
    use-on: not-win32  # 启动模式：any表示全平台、only-win32表示仅windows平台、not-win32表示除windows以外所有平台，never表示任何平台均不启用。
    sigint-exit: enable  # 收到 sigint 信号后退出 （Windows中可一半呢由ctrl+c触发）
    sigterm-exit: enable  # 收到 sigterm 信号后退出 （Windows中一般由系统欻）
    sighup-exit: enable  # 收到 sighup 信号后退出 
    sigquit-exit: enable  # 收到 sigquit 信号后退出（Windows中一般也ctrl+break触发）

win32-console:  # 控制台管理，比起处理信号量，在Windows平台使用控制台API更接近原生且合理。
    use-on: only-win32  # 启动方式：any或only-win32表示仅在windows平台启用。never/not-win32表示任何平台均不启用。
    ctrl-c-exit: enable  # 接收到ctrl+c是否退出
    ctrl-break-exit: enable  # 接收到ctrl+break是否退出
    console-close-recovery: disable  # 当用户关闭控制台后，是否启用一个新的临时的控制台输出日志（通常不建议，因为关闭控制台即意味着程序退出，只有5000ms的时间给程序进行清理操作。同时程序一般清理时间不会太久，可能在新控制台启用前就已经完成程序退出的所有准备）

server:  # 系统执行服务所需要的参数
    stop-wait-time: 10s  # 服务退出时，等待清理结束的最长时间。
```

### 一些值的设定

* `Name` 系统名称：
  * 第一优先级：配置文件中的设定
  * 第二优先级：命令行参数
  * 最后优先级：项目根目录下Name文件（仅第一行有效、不包含换行符）
* `SemanticVersioning` 语义化版本号：
  * 第一优先级：项目根目录下的`Version`文件，可以`v`或`V`开头，后接语义化版本号，否则视为不满足要求进入第二优先级。
  * 第二优先级：读取`git`中最新的`tag`，可以`v`或`V`开头，后接语义化版本号，否则视为不满足要求进入第二优先级。
    若最新的tag对应的提交不是当前的提交（且无`dev`标识），则添加`+dev-时间戳-当前版本提交的哈希值`到语义化版本号中。
  * 第三优先级：若无`tag`，但有当前提交记录的哈希值，则版本号为`0.0.0+dev-时间戳-当前版本提交的哈希值`。
  * 最后优先级：使用随机版本号，版本号为`0.0.0+dev-时间戳-随机值`，随机值在执行`go generate`时生成（位于文件`random_data.txt`中），`go build`后固定。
* `Version` 版本号：`SemanticVersioning`前添加`v`的字符串。

## 日后升级计划

1. 单元测试
2. GitHub Action
3. 对`Windows`服务的支持。

## 协议

本软件基于 [MIT LICENSE](/LICENSE) 发布。
了解更多关于 MIT LICENSE，请 [点击此处](https://mit-license.song-zh.com) 。
