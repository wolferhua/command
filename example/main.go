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