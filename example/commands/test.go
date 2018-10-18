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
