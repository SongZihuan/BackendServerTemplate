# Golang 后端服务 模板程序 - 介绍

使用 `Golang` 实现的后端服务模板程序。

## 介绍

本模板程序旨在实现一个 `Golang` 的后端服务，可以是 Web服务 也可以是其他。为了完成这个目的，我设计了一部分基础套件。

* 内置日志系统，可分级（debug、info、warn、error、panic），支持投递到标准输出、文件、日期切割的文件、自定义输出、多输出合并。输出到控制台时可以定义字体、颜色、高亮等美化细节。输出到控制台的数据精简提炼，而把详细数据输出到文件中。
* 内置断点标记系统，一旦遇到标记则会立马输出提示（生产环境下看将其关闭，但不用马上删掉代码中的标记），本功能有日志系统提供支持。
* 命令行参数读取，使用`cobra`和`pflag`共同解析，支持子命令、自动生成帮助文档、命令行自动补全。支持`version`、`license`、`report`、`check`等基础命令。
* 配置文件读取，使用`viper`解析配置文件，目前支持`yaml`和`json`两种格式。同时支持反向输出配置文件。
* 自动重启，当配置文件发生变更时自动重启服务（完全重启整个进程）。
* 多任务系统（多协程运行）或单任务系统，整个任务系统框架具备优雅推出机制。
* 服务运行独占线程，某些特殊服务运行时要固定在一个线程上（而不能让 `Goroutine` 随意切换）。
* 支持信号量程序退出机制，常用于`Linux`来实现优雅关闭。在`Windows`也能用使用，例如在控制台时`ctrl+c`会被`Golang`模拟成信号以达到跨平台目的。但通常可以在`Windows`上关闭信号量退出机制，转而使用下面的`Console`退出机制。
* 支持`win32API`的`ConsoleAPI`，读取`Console`关闭事件，实现退出机制（执行清理任务），仅能用于`Windows`。
* 服务注册、启动、关闭、重启、删除。主要用于Windows的服务管理，不过Linux也能使用。

## 协议与准则

本软件基于 [MIT LICENSE](/LICENSE) 发布，请务必阅读并遵守相关规定。 了解更多关于 MIT LICENSE，请 [点击此处](https://mit-license.song-zh.com) 。

我们欢迎各位贡献者、赞助商等参与项目的开发等各项工作。请务必阅读并遵守我们的 [行为准则](/CODE_OF_CONDUCT.md) 。

同时，欢迎各位贡献者阅读 [贡献者指南](CONTRIBUTING.md) ，有助于我们更好的交流和维护项目。

若您为项目做出了任何贡献，我们将会记录在 [贡献者名单](/CONTRIBUTORS.md) 上，若有任何遗漏可以随时通过 `Issue` 与我们取得联系。

## 文档

推荐阅读文档：[DeepWiki](https://deepwiki.com/SongZihuan/BackendServerTemplate) 。

**Note:** `DeepWiki`根据仓库的源码文件自动生成`WiKi`文档，现在被列为本项目参考文档之一和推荐阅读文档。

## 入口

入口文件在`src/cmd`下，目前分别有三个程序：

* `lion` 是使用控制单元的多服务演示程序。
* `tiger` 是直接运行服务的单服务演示程序。
* `monkey` 是服务安装演示程序。
* `giraffe` 是 `lion` 和 `tiger` 的合体。`giraffe lion` 相当于 `lion`，`giraffe tiger` 相当于 `tiger`，把两个程序集成于一起。

入口程序不直接包含太多的实际代码，真正的`main`函数位于`src\mainfunc`下。
程序的返回值代表程序的`Exit Code`。
一般采用`0`表示正确结束。

## 编译

### 必要文件

#### 编译时配置

编译时配置文件用于在编译期间就决定运行时程序的部分行为，并且不能在运行时被修改。该文件应放置在根目录下，命名为 `BUILD.yaml`。

```yaml
# 本项目存在共计4个应用程序，分别命名为：lion、tiger、monkey、giraffe。同时这些名字也被称为该程序的包名（作用见下文）。

# 默认配置，当某个项目（例如：lion）的配置不存在时，则应用此配置。
# 注意：仅当目标配置不存在时，会使用默认配置，当目标配置存在而某个配置项缺失时，配置项不会继承此处设置的默认配置，而是按系统内置的默认值除了。
default:
  name: ""
  auto-name: false  # 若设置为true时，不能 `name` 选项（也不能设置为空，必须缺失该选项）
  # name 和 auto-name 的关系“
  # name: 直接指定程序的名称
  # auto-name: 自动设置程序的名称（使用可执行文件的名称、若无法获取则使用包名）
  # 当关闭了auto-name且设置了name时（不为空），程序将直接使用该名字。
  # 当关闭了auto-name且没设置name时（缺失而非设置为空）使用包名。
  # 当关闭了auto-name且设置name为空时（非确实）是不允许的。
  # 当启用auto-name时且没设置name（缺失或者为空），将使用自动设置程序的名称（见上文）。
  # 当启用auto-name且设置了name（不为空）是不允许的。

lion:  # 包名
  name: lion
  env-prefix: LION  # 环境变量前缀。此时，获取环境变量 NAME 则实际上会获取 LION_NAME。

tiger:
  name: tiger

monkey:
  name: monkey
  service:  # 作为注册服务的相关配置
    name: TestService  # 服务名称 若为空则使用程序名称（例如，在此处为：monkey）
    display-name: Test Service  # 服务的显示名称（人类可读形式），若为空则和 name 一致
    describe: 一个简单的Go测试服务  # 服务的秒数
    
    # 参数来源
    #  no      无运行时参数（默认行为）
    #  install 在安装时指定参数，例如：monkey install a b c，其中 a b c 作为参数
    #  config  在本配置文件中的 argument-list 列表指定运行时参数
    argument-from: install
    # argument-from 为 config 时启用
    argument-list: []
    # 环境变量来源
    #  no      无运行时环境变量（默认行为）
    #  install 在安装时，根据 env-get-list 获取安装时的真实环境变量
    #  config  在本配置文件中的 env-set-list 中指定环境变量
    
    env-from: no
    # env-from 为 install 时启用
    env-get-list:
      - a  # 安装程序（monkey install）运行时，获取环境变量 a，并作为服务运行时的环境变量（例如安装时 a 的值为 b ，服务运行时也将得到环境变量 a 的值为 b）
      - c
    # env-from 为 config 时启用
    env-set-list:
      a: b  # 例如：映射环境变量 a 的值为 b
      c: d


giraffe:
  name: giraffe

```

**注意：本配置文件是编译时配置文件，在编译后配置文件包含在二进制文件中，此后可移除和修改文件系统上的配置文件而不影响编译好的程序。**

#### 构建信息文件

构建信息文件是通过 `gob` 编码的二进制数据文件，应放置在 `buildinfo/build.otd` 目录。此文件具有时效性，不应该推送到 `git` 仓库。

要生成此文件，需要使用 `go generate` 命令。请参见下文。

### 生成

每次编译器请先运行`go generate`命令，生成编译所需要的文件（例如构建信息文件）。

在项目根目录下执行 `go generate` 即可获取这些文件：

```shell
$ go generate ./...
```

其中，`./...`表示执行当前目录下的所有（含本目录和递归的子目录）`go:generate`命令。

### 编译

使用`go build`指令进行编译。如何包位于`github.com/SongZihuan/BackendServerTemplate/src/cmd/`下，例如`github.com/SongZihuan/BackendServerTemplate/src/cmd/lion`

所以，编译命令如下（以`lion`为例）：
```shell
$ go build github.com/SongZihuan/BackendServerTemplate/src/cmd/lion
```

若用于开发环境，可以按如下方式编译：
```shell
$ go build -o lion -trimpath -ldflags='-s -w -extldflags "-static"' github.com/SongZihuan/BackendServerTemplate/src/cmd/lion
```

其中：
 * `-o lion` 表示输出二进制文件的路径和名称
   * Windows： `-o lion.exe`
   * Linux/MacOS：`-o lion`
 * `-trimpath` 表示擦除二进制文件中关于源码目录的信息，可以压缩二进制文件体积和包含隐私。
 * `-ldflags='-s -w'` 是传递给链接器的参数。
   * `-s` 表示去掉符号和调试信息，可以减少二进制文件大小，同时增加反编译的难度。
   * `-w` 表示去掉`DWARF`调试信息，可以减少二进制文件大小，同时增加反编译的难度。
   * `-extldflags "-static"` 是传递给外部链接器的参数。
     * `-static` 表示外部链接器生成静态链接文件（实际上我几乎从真正看过这个参数的作用，因为`go`本身就会优先以静态形式链接库文件，这个参数可能和`cgo`搭配更合适）。

## 运行
执行编译好地可执行文件即可。
**注意编译时选择的目标平台要与运行平台一致。**

### 运行参数

接下来所列出来的指令，是`lion`、`tiger`和`monkey`所共有的子命令和标志。
而`giraffe lion`的参数和`lion`相同，`giraffe tiger`的参数和`tiger`相同，不再分开介绍。

下列为公共子命令，以`lion`为例子：

* `lion` 无子命令时，表示直接运行该程序。
  * 注意：直接调用`monkey`运行代码是具有未知性的，应该使用`monkey start`。
* `lion help` 查看`lion`的帮助文档。
  * `lion help <subcommand>` 查看子命令的帮助文档，例如：`lion help version`。
* `lion version` 输出详细的版本号。
* `lion license` 输出授权协议（位于项目根目录的`LICENSE`）。
* `lion report` 输出程序反馈、报告方式文档。
* `lion check` 检查并输出配置文件的检查接管。
* `completion` 与`shell`和`powershell`实现自动补全有关。

下列为公共标志，以`lion`为例子：

* `lion -n xxx`，`lion --name xxx` 设置程序运行的名称，需要一个参数 **（全局参数，任何子命令均具有此参数）**。
* `lion -h`，`lion --help` 查看帮助文档 **（全局参数，任何子命令均具有此参数）**。
* `lion -c xxx`，`lion --config xxx` 配置文件位置（可提供一个字符串参数作为路径指向配置文件，默认值为：`config.yaml`）。
* `lion -o xxx`、`lion --output-config xxx` 反向输出配置文件（默认不输出，若提供一个路径则会输出到该路径所指处）。未知配置项不会输出，未设定配置项则以默认值（若有）输出。**前提：配置文件被正确导入和识别。**

以下标志为`tiger`和`lion`共有，`monkey`没有的：

* `lion --auto-reload` 自动重载（监测配置文件，若其变化则重启服务系统）。该功能应该仅仅用在开发时。

以下为`check`子命令的标志：

* `lion check -c xxx`，`lion check --config xxx` 配置文件位置（可提供一个字符串参数作为路径指向配置文件，默认值为：`config.yaml`）。
* `lion check -o xxx`、`lion check --output-config xxx` 反向输出配置文件（默认不输出，若提供一个路径则会输出到该路径所指处）。未知配置项不会输出，未设定配置项则以默认值（若有）输出。**前提：配置文件被正确导入和识别。**

以下为`version`子命令的标志：

* `lion version -s` `lion version --short` 仅输出短的版本号。

特定于`monkey`的子命令（如：`install`、`start`等）可参考：[后端服务](#后台服务)

### 运行时配置文件

**注意：此配置文件为运行时配置文件，即编译后的运行时阶段才从指定文件路径中读取。与后面提及的服务注册所用的`SERVICE.yaml`编译时配置文件不同。**

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

    warn-writer:  # debug、info、warning的输出器
      - type: stander  # 输出类型（stander - 标准输出）
        format: console-try-pretty  # 格式：console（比file模式缩减内容，用于控制台快速查看信息）、console-pretty（在 console 基础上开启 ANSI 转义序列）、console-try-pretty（类似于 console-pretty，但是会检测终端是否支持ANSI，若不支持则不启用ANSI）
        output-path: stdout  # 输出位置（对于type为stander来说，只能为stdout和stderr）
      - type: file  # 输出类型（file - 文件）
        format: file  # 格式：文件类型（比console多更多复杂的记录）
        output-path: /path/to/warn.log  # 存储日志的文件（若文件或路径不存在会尝试创建）
      - type: date-file  # 输出类型（date-file - 按日期切割日志）
        format: file  # 同上
        output-path: /path/to/log  # 存储日志的文件夹（若不存在会尝试创建）
        file-prefix: warn-date  # 日志文件的文件名前缀，例：warn-date.2025-05-10.log
      - type: date-file  # 同上
        format: json  # json格式，一行json为一条日志，一般用于机器读取
        output-path: /path/to/log  # 同时
        file-prefix: warn-date-machine  # 同时

    err-writer:  # error、panic 的输出器
        - type: stander
          format: try-pretty
          output-path: stderr
        - type: file
          format: file
          output-path: C:\Users\songz\Code\GoProject\BackendServerTemplate\test_self\error.log
        - type: date-file
          format: file
          output-path: C:\Users\songz\Code\GoProject\BackendServerTemplate\test_self
          file-prefix: error-date
        - type: date-file
          format: json
          output-path: C:\Users\songz\Code\GoProject\BackendServerTemplate\test_self
          file-prefix: error-date-machine

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
  name: Jim
  example1:
    stop-wait-time: 10s  # 服务退出时，等待清理结束的最长时间，若超过可能会因系统强制下线而导致一些问题。
    startup-wait-time: 3s  # 服务启动需要的等待时间（若在该时间内未响应服务启动成功则会产生 timeout 警告。一般情况下 3s 即可）
    lock-thread: disable  # 是否开启独占进程（默认不开启）
  example2:
    stop-wait-time: 10s
    startup-wait-time: 3s
    lock-thread: disable
  example3:
    stop-wait-time: 10s
    startup-wait-time: 3s
    lock-thread: disable
  controller:
    stop-wait-time: 10s
    stop-wait-time-use-specified-value: disable  # 若启用，控制器真实的清理等待时间只等于 stop-wait-time ，若不启用则等于 stop-wait-time + 任务中 stop-wait-time 最大的那一个。
    startup-wait-time: 3s
    startup-wait-time-use-specified-value: disable  # 若启用，控制器真实的启动等待时间只等于 startup-wait-time ，若不启用则等于 startup-wait-time + 任务中 startup-wait-time 最大的那一个。
    lock-thread: disable
```

### 一些值的设定

* `Name` 系统名称：
  * 第一优先级：配置文件中的设定
  * 第二优先级：命令行参数
  * 最后优先级：项目根目录下Name文件（仅第一行有效、不包含换行符）
* `SemanticVersion` 语义化版本号：
  * 第一优先级：项目根目录下的`Version`文件，可以`v`或`V`开头，后接语义化版本号，否则视为不满足要求进入第二优先级。
  * 第二优先级：读取`git`中最新的`tag`，可以`v`或`V`开头，后接语义化版本号，否则视为不满足要求进入第二优先级。
    若最新的tag对应的提交不是当前的提交（且无`dev`标识），则添加`+dev.时间戳.当前版本提交的哈希值`到语义化版本号中。
  * 第三优先级：若无`tag`，但有当前提交记录的哈希值，则版本号为`0.0.0+dev.时间戳.当前版本提交的哈希值`。
  * 最后优先级：使用随机版本号，版本号为`0.0.0+dev.时间戳.随机值`，随机值在执行`go generate`时生成（位于文件`random_data.txt`中），`go build`后固定。
* `Version` 版本号：`SemanticVersion`前添加`v`的字符串。

## 后台服务

虽然`lion`和`tiger`也可以作为后台服务，但是我使用了`monkey`进行了更高层次的抽象，使得在`Windows`和`Linux`上可以安装服务程序。

后台服务采用Go的第三方库`github.com/kardianos/service`实现，主要目的是实现`Windows`上的服务注册。
但是理论上来说，`MacOS`和`Linux`（`systemd`）也能使用。
不过，在`Linux`上注册服务，可能自己编辑`systemd`配置文件，或者使用宝塔等辅助面板会更为灵活。

### 安装

```shell
$ monkey install <命令行参数列表>
```

使用此命令可以在`Windows`中或`Linux`中注册一个服务.

注意：安装后可执行程序`monkey`仍需保留在原来位置，不可移动。

### 卸载

```shell
$ monkey uninstall
```

或者

```shell
$ monkey remove
```

### 启动

```shell
$ monkey start
```

启动不需要指定命令行参数，命令行参数在`install`时即确定。

### 停止

```shell
$ monkey stop
```

### 重启

```shell
$ monkey restart
```
