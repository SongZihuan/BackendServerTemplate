# Example: 例子

此处是一份配置文件，可以在项目根目录处运行。

## 文件说明

- [README.md](./README.md) 自述文件。
- [config.example.yaml](./config.example.yaml) 配置文件举例。
- [.gitignore](./.gitignore)`.gitignore` 配置`git`忽略文件，运行时会在本文件夹生成一些文件（例如：日志）。

## 使用

以`lion`距离，假设编译出来的可执行文件位于`/path/of/lion.exe`，则：

```shell
$ cd /path/of/the/project # 移动到项目根目录
$ /path/of/lion.exe --config ./example/config.example.yaml # /path/of/lion.exe 替换为可执行文件的路径
```
