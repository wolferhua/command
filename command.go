package command

import (
	"flag"
	"os"
	"sync"
)

//命令行列表
var commands = map[string]CmdInterface{}

//命令行列表锁
var mu = sync.Mutex{}

//程序初始化
func Init() {
	Registry(
		&Help{},
	)
}

//注册命令
func Registry(cs ...CmdInterface) {
	mu.Lock()
	defer mu.Unlock()
	for _, c := range cs {
		name := c.GetName()
		if len(name) > 0 {
			_, ok := commands[name]
			if ok {
				Error("command " + name + " repeat")
				continue
			}
			commands[name] = c
		}
	}

}

//获取命令行
func Get(name string) (c CmdInterface, ok bool) {
	c, ok = commands[name]
	return
}

//获取所有数据
func GetAll() map[string]CmdInterface {
	return commands
}


//启动command
func Run() bool {
	Init()
	argsLen := len(os.Args)
	if argsLen > 1 {
		name := os.Args[1]
		c, ok := Get(name)
		if ok {
			//参数设置
			c.SetFlags()
			flag.CommandLine.Parse(os.Args[2:])
			//初始化
			c.Init()
			//运行
			c.Run()
		}
		return ok
	}
	return false
}
