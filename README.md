# command
go command library

# 项目描述

go命令行管理工具，go 原生提供了flag管理命令行。使用起来诸多不便，所以整理如下框架实现。

# 使用入门

### 第〇步： 获取command

```shell
go get -u github.com/wolferhua/command
```

### 第一步： 建立命令行文件目录

```shell
mkdir -p project/commands
```

### 第二步： 创建命令文件 commands/test.go

```go
package commands

import (
	"fmt"
	"github.com/wolferhua/command"
)

type Test struct {
	command.Cmd
	B bool
}

//设置名称
func (c *Test) GetName() string {
	return "test"
}

//设置描述
func (c *Test) GetDesc() string {
	return "test demo"
}
//运行
func (c *Test) Run() {
	fmt.Println("test is run")
	fmt.Println(c.B)
}

//参数设置
func (c *Test) SetFlags() {
	c.BoolVar(&c.B, "b", false, "bool test")
}

```


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
### 第四步： main中运行

```go
package main

import (
	"github.com/wolferhua/command"
	_ "github.com/wolferhua/command/example/commands" //注册commands
)

func main() {
	if command.Run() {
		//command 运行了。可以退出程序
	}else{
		//没有运行，可以运行其他操作
	}
}
```


### 第五步： 运行

**参数运行**
```bash
go run example/main.go test -b
```

```bash
# test is run
# true //接收数据
```

**不带参数运行**

```bash
go run example/main.go test 
```

```bash
# test is run
# false //没有数据
```

