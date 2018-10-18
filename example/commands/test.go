package commands

import (
	"fmt"
	"github.com/wolferhua/command"
)

type Test struct {
	command.Cmd
	B bool
}

func (c *Test) GetName() string {
	return "test"
}

func (c *Test) GetDesc() string {
	return "test demo"
}
func (c *Test) Run() {
	fmt.Println("test is run")
	fmt.Println(c.B)
}

func (c *Test) SetFlags() {
	c.BoolVar(&c.B, "b", false, "bool test")
}
