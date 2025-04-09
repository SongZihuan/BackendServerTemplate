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

使用请先在项目根目录执行`go:generate ./...`，随后执行`go build github.com/SongZihuan/BackendServerTemplate/src/cmd/<lionv1/tigerv1>`进行正常编译即可。
**注意：把`<lionv1/tigerv1>`替换成你具体想编译的包名，最终指令例如：go build github.com/SongZihuan/BackendServerTemplate/src/cmd/lionv1**

日后支持：

* 添加编译参数，允许编译时注入编译时间和`git commit id`。

### 运行
执行编译好的可执行文件即可。具体命令行参数可参见上文。注意编译时选择的目标平台要与运行平台一致。

## 协议
本软件基于 [MIT LICENSE](/LICENSE) 发布。
了解更多关于 MIT LICENSE , 请 [点击此处](https://mit-license.song-zh.com) 。
