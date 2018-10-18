# command
go command library

# 项目描述

go命令行管理工具，go 原生提供了flag管理命令行。使用起来诸多不便，所以整理如下框架实现。

# 使用入门

### 第一步： 建立命令行文件目录

```shell
mkdir commands
```

### 第二步： 创建命令文件 commands/test.go


### 第三步： 创建命令注册文件 commands/commands.go

```go
package commands

import "github.com/wolferhua/command"

func init() {
	command.Registry(
		&Test{},
	)
}
```



