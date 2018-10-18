package commands

import (
	"fmt"
	"github.com/wolferhua/command"
)

type Test struct {
	command.Cmd
	b bool
}

func (c *Test) GetName() string {
	return "test"
}

func (c *Test) GetDesc() string {
	return "test demo"
}
func (c *Test) Run() {
	fmt.Println("test is run")
}

func (c *Test) SetFlags() {
	c.BoolVar(&c.b, "b", false, "bool test")
}
